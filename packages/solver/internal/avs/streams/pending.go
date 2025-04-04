package streams

import (
	"context"
	"solver/internal/redis"

	redisv8 "github.com/go-redis/redis/v8"
)

func GetPendingMessages(ctx context.Context, consumerGroup string) (*redisv8.XPending, error) {
	return redis.StreamsRedis.XPending(ctx, IntentStream, consumerGroup).Result()
}
