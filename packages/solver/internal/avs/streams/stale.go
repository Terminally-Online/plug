package streams

import (
	"context"
	"solver/internal/redis"
	"time"

	redisv8 "github.com/go-redis/redis/v8"
)

func ClaimStaleMessages(ctx context.Context, consumerGroup, consumer string, minIdleTime time.Duration, start, end string, count int64) ([]redisv8.XMessage, error) {
	return redis.StreamsRedis.XClaim(ctx, &redisv8.XClaimArgs{
		Stream:   IntentStream,
		Group:    consumerGroup,
		Consumer: consumer,
		MinIdle:  minIdleTime,
		Messages: []string{start, end},
	}).Result()
}
