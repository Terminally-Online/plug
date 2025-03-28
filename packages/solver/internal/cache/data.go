package cache

import (
	"context"
	"encoding/json"
	"solver/internal/redis"
	"time"
)

type CacheData struct {
	Value     json.RawMessage `json:"value"`
	CreatedAt time.Time       `json:"created_at"`
}

func WithCache[T any](key string, options []CacheOption, fn func() (T, error)) (T, error) {
	opts := defaultCacheOptions()

	for _, opt := range options {
		opt(opts)
	}

	data, err := redis.CacheRedis.Get(context.Background(), key).Bytes()

	if err == nil {
		var cachedData CacheData
		if err := json.Unmarshal(data, &cachedData); err == nil {
			var result T
			if err := json.Unmarshal(cachedData.Value, &result); err == nil {
				if time.Since(cachedData.CreatedAt) <= opts.duration {
					return result, nil
				}

				if opts.useStale {

					go func() {
						updateCache(context.Background(), key, fn, opts)
					}()
					return result, nil
				}
			}
		}

	}

	return updateCache(context.Background(), key, fn, opts)
}

func updateCache[T any](ctx context.Context, key string, fn func() (T, error), opts *CacheOptions) (T, error) {
	result, err := fn()
	if err != nil {
		var zero T
		return zero, err
	}

	valueData, err := json.Marshal(result)
	if err == nil {
		cacheData := CacheData{
			Value:     valueData,
			CreatedAt: time.Now(),
		}

		cacheBytes, err := json.Marshal(cacheData)
		if err == nil {
			expiration := opts.duration
			if opts.useStale {
				expiration = opts.duration + opts.staleBuffer
			}
			_ = redis.CacheRedis.Set(ctx, key, cacheBytes, expiration).Err()
		}
	}

	return result, nil
}
