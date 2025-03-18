package cache

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Redis client configuration
var (
	Redis = redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "",
		DB:          0,
		DialTimeout: 100 * time.Millisecond,
		ReadTimeout: 100 * time.Millisecond,
	})
	defaultCachePeriod = 5 * time.Minute
)

type cacheData struct {
	Value     json.RawMessage `json:"value"`
	CreatedAt time.Time       `json:"created_at"`
}

// GenerateCacheKey creates a cache key from various data types
func GenerateCacheKey(v interface{}) (string, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	case int, int64, uint64, float64, bool:
		return fmt.Sprintf("%v", v), nil
	default:
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			return "", fmt.Errorf("failed to marshal value: %w", err)
		}

		hash := sha256.Sum256(jsonBytes)
		return hex.EncodeToString(hash[:32]), nil
	}
}

/**
 * WithCache fetches data from cache or executes the function and caches the result.
 * If a useStale option is set, stale data will be returned while the cache is updated in the background.
 * The staleBuffer option sets the additional time to keep stale data for immediate use before redis cleanup.
 * Example usage:
 * result, err := cache.WithCache("key", cache.WithOptions(
 *     cache.WithDuration(5*time.Minute),
 *     cache.WithStaleData(true)
 *   ),
 *   func() (MyType, error) {}
 * )
 */
func WithCache[T any](key string, options []CacheOption, fn func() (T, error)) (T, error) {
	opts := defaultCacheOptions()
	// Apply options if provided
	for _, opt := range options {
		opt(opts)
	}

	// Try to get from cache
	data, err := Redis.Get(context.Background(), key).Bytes()

	// Handle cache hit
	if err == nil {
		var cachedData cacheData
		if err := json.Unmarshal(data, &cachedData); err == nil {
			var result T
			if err := json.Unmarshal(cachedData.Value, &result); err == nil {
				// Check if the data is fresh
				if time.Since(cachedData.CreatedAt) <= opts.duration {
					return result, nil
				}

				// Data is stale but we can use it
				if opts.useStale {
					// Use stale data and update in background
					go func() {
						updateCache(context.Background(), key, fn, opts)
					}()
					return result, nil
				}
			}
		}
		// If unmarshal fails or data is stale, continue to fetch fresh data
	}

	// No valid cache or stale data not requested - do synchronous update
	return updateCache(context.Background(), key, fn, opts)
}

// updateCache executes the function and updates the cache
func updateCache[T any](ctx context.Context, key string, fn func() (T, error), opts *cacheOptions) (T, error) {
	result, err := fn()
	if err != nil {
		var zero T
		return zero, err
	}

	// Cache the result
	valueData, err := json.Marshal(result)
	if err == nil {
		cacheData := cacheData{
			Value:     valueData,
			CreatedAt: time.Now(),
		}

		cacheBytes, err := json.Marshal(cacheData)
		if err == nil {
			// Set the cache with the stale buffer added to expiration
			expiration := opts.duration
			if opts.useStale {
				expiration = opts.duration + opts.staleBuffer
			}
			_ = Redis.Set(ctx, key, cacheBytes, expiration).Err()
		}
	}

	return result, nil
}

// CacheOptions holds configuration for cache behavior
type cacheOptions struct {
	duration    time.Duration
	useStale    bool
	staleBuffer time.Duration
}

// CacheOption is a function that modifies cache options
type CacheOption func(*cacheOptions)

func defaultCacheOptions() *cacheOptions {
	return &cacheOptions{
		duration:    defaultCachePeriod,
		useStale:    false,
		staleBuffer: 10 * time.Minute,
	}
}

// Wraps options for a cleaner call and implementation
func WithOptions(options ...CacheOption) []CacheOption {
	return options
}

// WithDuration sets the cache duration for a cache key
func WithDuration(duration time.Duration) CacheOption {
	return func(o *cacheOptions) {
		o.duration = duration
	}
}

// WithStaleData allows using stale cache data for a cache key
func WithStaleData(useStale bool) CacheOption {
	return func(o *cacheOptions) {
		o.useStale = useStale
	}
}

// WithStaleBuffer sets the additional time to keep stale data for the ability to use it immediately and update in background
func WithStaleBuffer(buffer time.Duration) CacheOption {
	return func(o *cacheOptions) {
		o.staleBuffer = buffer
	}
}
