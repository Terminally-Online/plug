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

type UserRequests struct {
	count       atomic.Int64
	limit       int
	windowStart atomic.Int64 // Unix timestamp
	isRevoked   bool
}

type RateLimiter struct {
	// Still need a mutex for map operations, but not for counter operations
	requests map[string]*UserRequests
	mu       sync.RWMutex
}

func NewRateLimiter() *RateLimiter {
	limiter := &RateLimiter{
		requests: make(map[string]*UserRequests),
	}

	return limiter
}

func (rl *RateLimiter) Allow(apiKey string) bool {
	now := time.Now().Unix()

	// Read lock for lookup only
	rl.mu.RLock()
	userReq, exists := rl.requests[apiKey]
	rl.mu.RUnlock()

	if !exists {
		rl.mu.Lock()

		// Check again with write lock, if it still doesn't exist, create a new entry
		if userReq, exists = rl.requests[apiKey]; !exists {
			dbApiKey, err := lookupApiKey(apiKey)
			if err != nil {
				rl.mu.Unlock()
				return false
			}
			userReq = &UserRequests{}
			userReq.windowStart.Store(now)
			userReq.count.Store(1)
			userReq.limit = dbApiKey.RateLimit

			if dbApiKey.RevokedAt != nil {
				userReq.isRevoked = true
				rl.mu.Unlock()
				return false
			}

			userReq.isRevoked = false
			rl.requests[apiKey] = userReq
			rl.mu.Unlock()
			return true
		}

		rl.mu.Unlock()
	}

	// Check if window expired and reset if needed
	windowStart := userReq.windowStart.Load()
	if now-windowStart >= WINDOW_SECONDS {
		dbApiKey, err := lookupApiKey(apiKey)
		if err != nil {
			rl.requests[apiKey] = nil
			return false
		}
		userReq.windowStart.Store(now)
		userReq.count.Store(1)
		userReq.limit = dbApiKey.RateLimit
		userReq.isRevoked = dbApiKey.RevokedAt != nil
		if userReq.isRevoked {
			return false
		}
		return true
	}

	// Check rate limit and increment atomically
	currentCount := userReq.count.Load()
	if currentCount >= LIMIT {
		return false
	}

	// Try to increment using CompareAndSwap to handle race conditions
	// If another request incremented past limit concurrently, this will fail
	return userReq.count.CompareAndSwap(currentCount, currentCount+1)
}

func lookupApiKey(apiKey string) (*models.ApiKey, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("missing api key")
	}

	dbApiKey := database.DB.Where("key = ?", apiKey).First(&models.ApiKey{})
	if dbApiKey.Error != nil {
		return nil, fmt.Errorf("invalid api key")
	}

	if dbApiKey.RevokedAt {
		return nil, fmt.Errorf("revoked api key")
	}

	return dbApiKey, nil
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-Api-Key")
		if key == "" {
			key = r.RemoteAddr
		}

		if !rl.Allow(key) {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Rate limit exceeded"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) ApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-Api-Key")
		if apiKey == "" {
			http.Error(w, "missing api key", http.StatusUnauthorized)
			return
		}

		dbApiKey := database.DB.Where("key = ?", apiKey).First(&models.ApiKey{})
		if dbApiKey.Error != nil {
			http.Error(w, "invalid api key", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) AdminApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Api-Key") != h.appApiKey {
			http.Error(w, "invalid api key", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
