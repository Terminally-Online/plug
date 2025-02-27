package utils

import (
	"sync"
	"time"
)

type cacheEntry struct {
	Data   interface{}
	Expiry time.Time
}

// Note that sync.Map is thread safe and does not require a mutex as it's handled internally via atomic operations
var (
	cache              sync.Map
	defaultCachePeriod = 5 * time.Minute
)

// WithCache is a generic caching function that caches the result of fn
// If useStale is true, it will return stale data (if available) and update the cache in the background
func WithCache[T any](key string, duration []time.Duration, useStale bool, fn func() (T, error)) (T, error) {
	cachePeriod := defaultCachePeriod
	if len(duration) > 0 {
		cachePeriod = duration[0]
	}

	if cached, ok := cache.Load(key); ok {
		entry := cached.(cacheEntry)

		// If the cache is still valid, return it
		if entry.Expiry.After(time.Now()) {
			return entry.Data.(T), nil
		}

		// If useStale is true and we have stale data, return it and update in background
		if useStale {
			go func() {
				result, err := fn()
				if err != nil {
					return
				}

				cache.Store(key, cacheEntry{
					Data:   result,
					Expiry: time.Now().Add(cachePeriod),
				})
			}()

			return entry.Data.(T), nil
		}

		cache.Delete(key)
	}

	// No cache or stale data not requested - do synchronous update
	result, err := fn()
	if err != nil {
		var zero T
		return zero, err
	}

	cache.Store(key, cacheEntry{
		Data:   result,
		Expiry: time.Now().Add(cachePeriod),
	})

	return result, nil
}
