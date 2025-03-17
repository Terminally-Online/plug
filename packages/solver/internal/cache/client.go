package cache

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"solver/internal/actions"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	Redis = redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "",
		DB:          0,
		DialTimeout: 100 * time.Millisecond,
		ReadTimeout: 100 * time.Millisecond,
	})
	CacheDuration = 5 * time.Minute
)

func cacheKey[T any](v T) (string, error) {
	switch v := any(v).(type) {
	case string:
		return v, nil
	case int:
		return fmt.Sprintf("%d", v), nil
	case int64:
		return fmt.Sprintf("%d", v), nil
	case uint64:
		return fmt.Sprintf("%d", v), nil
	case float64:
		return fmt.Sprintf("%f", v), nil
	case bool:
		return fmt.Sprintf("%t", v), nil
	default:
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			return "", fmt.Errorf("failed to marshal value: %w", err)
		}

		hash := sha256.Sum256(jsonBytes)
		return hex.EncodeToString(hash[:32]), nil
	}
}

func Get[TReturn map[string]actions.ProtocolSchema, TParams any](params TParams) (TReturn, error) {
	cacheKey, err := cacheKey(params)
	if err != nil {
		return nil, fmt.Errorf("failed to generate cache key: %w", err)
	}

	cachedData, err := Redis.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var result TReturn
		if err := json.NewDecoder(strings.NewReader(cachedData)).Decode(&result); err != nil {
			return nil, fmt.Errorf("failed to decode cache: %w", err)
		}
		return result, nil
	}
	if err != redis.Nil {
		return nil, fmt.Errorf("cache error: %w", err)
	}

	return nil, nil
}

func Set[TData map[string]actions.ProtocolSchema, TParams any](params TParams, data TData) error {
	cacheKey, err := cacheKey(params)
	if err != nil {
		return fmt.Errorf("failed to generate cache key: %w", err)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal response for cache: %w", err)
	}

	return Redis.Set(context.Background(), cacheKey, jsonData, CacheDuration).Err()
}
