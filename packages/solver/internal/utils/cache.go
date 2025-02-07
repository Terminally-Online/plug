package utils

import (
	"sync"
	"time"
)

type cacheEntry struct {
	Data   interface{}
	Expiry time.Time
}

var (
	cache              sync.Map
	defaultCachePeriod = 5 * time.Minute
)

// WithCache is a generic caching function that caches the result of fn
func WithCache[T any](key string, duration []time.Duration, fn func() (T, error)) (T, error) {
	cachePeriod := defaultCachePeriod
	if len(duration) > 0 {
		cachePeriod = duration[0]
	}

	if cached, ok := cache.Load(key); ok {
		if cache := cached.(cacheEntry); cache.Expiry.After(time.Now()) {
			return cache.Data.(T), nil
		}
		cache.Delete(key)
	}

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
