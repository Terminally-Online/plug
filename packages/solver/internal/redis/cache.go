package redis

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"context"

	"github.com/go-redis/redis/v8"
)

var (
	CacheDatabase    = 0
	CacheDialTimeout = 100 * time.Millisecond
	CacheReadTimeout = 100 * time.Millisecond
	CacheRedis       = redis.NewClient(&redis.Options{
		Addr:        RedisAddress,
		Password:    os.Getenv("REDIS_PASSWORD"),
		DB:          CacheDatabase,
		DialTimeout: CacheDialTimeout,
		ReadTimeout: CacheReadTimeout,
	})
)

func GenerateCacheKey(v any) (string, error) {
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

// ClearCache flushes all keys in the cache database
func ClearCache(ctx context.Context) error {
	return CacheRedis.FlushDB(ctx).Err()
}
