package streams

import (
	"context"
	"solver/internal/redis"
)

// NOTE: As of now there is no recourse for an acknowledged, but unresponsive action.

func AckMessage(ctx context.Context, consumerGroup, messageID, stream string) error {
	return redis.StreamsRedis.XAck(ctx, stream, consumerGroup, messageID).Err()
}
