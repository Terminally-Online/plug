package middleware

import (
	"fmt"
	"net/http"
	"solver/internal/database"
	"solver/internal/database/models"
	"sync"
	"sync/atomic"
	"time"
)

const (
	WINDOW_SECONDS = 60
)

type KeyCache struct {
	key         *models.ApiKey
	count       atomic.Int64
	windowStart atomic.Int64
}

type RateLimitInfo struct {
	Used      int64     `json:"used"`
	Limit     int       `json:"limit"`
	Reset     time.Time `json:"reset"`
	Remaining int64     `json:"remaining"`
}

type KeyRateLimiter struct {
	keys map[string]*KeyCache
	mu   sync.RWMutex // mutex for map operations
}

// GetLimit returns rate limit information for a given API key
func (rl *KeyRateLimiter) GetLimit(apiKeyID string) (RateLimitInfo, bool) {
	rl.mu.RLock()
	defer rl.mu.RUnlock()
	
	for _, cache := range rl.keys {
		if cache.key.Id == apiKeyID {
			currentCount := cache.count.Load()
			resetTime := time.Unix(cache.windowStart.Load()+WINDOW_SECONDS, 0)
			return RateLimitInfo{
				Used:      currentCount,
				Limit:     cache.key.RateLimit,
				Reset:     resetTime,
				Remaining: int64(cache.key.RateLimit) - currentCount,
			}, true
		}
	}
	
	return RateLimitInfo{}, false
}

func NewKeyRateLimiter() *KeyRateLimiter {
	return &KeyRateLimiter{
		keys: make(map[string]*KeyCache),
	}
}

/*
Allow does a few things to facilitate the rate limiting of an API key.
1. We check if the key currently exists in the rate limit cache. If it doesn't, we create a new entry in the cache.
2. If it does exist, we check if the window has expired. If it has, we reset the window and count, and refetch the key details from the database.
3. We check if the rate limit has been reached, and if not, we increment the count atomically.
4. If the rate limit has been reached, we return false, otherwise return true.
*/
func (rl *KeyRateLimiter) Allow(apiKey string) (keyModel models.ApiKey, limitInfo RateLimitInfo, statusCode int, err error) {
	now := time.Now().Unix()

	rl.mu.RLock()
	cachedKey, exists := rl.keys[apiKey]
	rl.mu.RUnlock()

	// If the key does not exist in cache yet, look it up and create a new entry
	if !exists {
		rl.mu.Lock()
		defer rl.mu.Unlock()

		// Check again with write lock, if it still doesn't exist, create a new entry
		if cachedKey, exists = rl.keys[apiKey]; !exists {
			dbApiKey, statusCode, err := lookupApiKey(apiKey)
			if err != nil {
				return models.ApiKey{}, RateLimitInfo{}, statusCode, err
			}

			cachedKey = &KeyCache{
				key: dbApiKey,
			}
			cachedKey.windowStart.Store(now)
			cachedKey.count.Store(1)
			rl.keys[apiKey] = cachedKey

			resetTime := time.Unix(cachedKey.windowStart.Load()+WINDOW_SECONDS, 0)
			currentCount := cachedKey.count.Load()
			limitInfo = RateLimitInfo{
				Used:      currentCount,
				Limit:     cachedKey.key.RateLimit,
				Reset:     resetTime,
				Remaining: int64(cachedKey.key.RateLimit) - currentCount,
			}

			return *cachedKey.key, limitInfo, http.StatusOK, nil
		}
	}

	// Check if window expired and reset count and window if it has
	windowStart := cachedKey.windowStart.Load()
	if now-windowStart >= WINDOW_SECONDS {
		rl.mu.Lock()

		dbApiKey, statusCode, err := lookupApiKey(apiKey)
		if err != nil {
			delete(rl.keys, apiKey)
			rl.mu.Unlock()
			return models.ApiKey{}, RateLimitInfo{}, statusCode, err
		}

		cachedKey.key = dbApiKey
		cachedKey.windowStart.Store(now)
		cachedKey.count.Store(1)

		resetTime := time.Unix(cachedKey.windowStart.Load()+WINDOW_SECONDS, 0)
		currentCount := cachedKey.count.Load()
		limitInfo = RateLimitInfo{
			Used:      currentCount,
			Limit:     cachedKey.key.RateLimit,
			Reset:     resetTime,
			Remaining: int64(cachedKey.key.RateLimit) - currentCount,
		}

		rl.mu.Unlock()
		return *cachedKey.key, limitInfo, http.StatusOK, nil
	}

	currentCount := cachedKey.count.Load()
	resetTime := time.Unix(cachedKey.windowStart.Load()+WINDOW_SECONDS, 0)
	limitInfo = RateLimitInfo{
		Used:      currentCount,
		Limit:     cachedKey.key.RateLimit,
		Reset:     resetTime,
		Remaining: int64(cachedKey.key.RateLimit) - currentCount,
	}

	if currentCount >= int64(cachedKey.key.RateLimit) {
		return *cachedKey.key, limitInfo, http.StatusTooManyRequests, fmt.Errorf("rate limit exceeded")
	}

	// Update the Used count after incrementing
	newCount := cachedKey.count.Add(1)
	limitInfo.Used = newCount
	limitInfo.Remaining = int64(cachedKey.key.RateLimit) - newCount
	
	return *cachedKey.key, limitInfo, http.StatusOK, nil
}

func lookupApiKey(apiKey string) (dbKey *models.ApiKey, statusCode int, err error) {
	if apiKey == "" {
		return nil, http.StatusUnauthorized, fmt.Errorf("missing api key")
	}

	dbApiKey := &models.ApiKey{}
	result := database.DB.Unscoped().Where("key = ?", apiKey).First(dbApiKey)
	if result.Error != nil {
		return nil, http.StatusUnauthorized, fmt.Errorf("invalid api key")
	}

	if !dbApiKey.DeletedAt.Time.IsZero() {
		return nil, http.StatusUnauthorized, fmt.Errorf("revoked api key")
	}

	return dbApiKey, http.StatusOK, nil
}

func (h *Middleware) AdminApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-Api-Key")

		dbKey, _, statusCode, err := h.apiKeyLimiter.Allow(apiKey)
		if err != nil {
			http.Error(w, err.Error(), statusCode)
			return
		}

		if dbKey.Role != "admin" {
			http.Error(w, "invalid api key role", http.StatusUnauthorized)
			return
		}

		r.Header.Set("X-Api-Key-Id", dbKey.Id)

		next.ServeHTTP(w, r)
	})
}

func (h *Middleware) ApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-Api-Key")
		dbKey, limitInfo, statusCode, err := h.apiKeyLimiter.Allow(apiKey)
		if err != nil {
			http.Error(w, err.Error(), statusCode)
			return
		}

		r.Header.Set("X-Api-Key-Id", dbKey.Id)
		w.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", limitInfo.Limit))
		w.Header().Set("X-RateLimit-Used", fmt.Sprintf("%d", limitInfo.Used))
		w.Header().Set("X-RateLimit-Reset", fmt.Sprintf("%d", limitInfo.Reset.Unix()))
		w.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", limitInfo.Remaining))

		next.ServeHTTP(w, r)
	})
}