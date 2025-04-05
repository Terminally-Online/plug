package streams

import (
	"context"
	"encoding/json"
	"solver/internal/redis"
	"time"

	redisv8 "github.com/go-redis/redis/v8"
)

func Publish(ctx context.Context, intentId string, data map[string]any) error {
	message, err := json.Marshal(data)
	if err != nil {
		return err
	}
	
	var chainID uint64
	if chainIDValue, ok := data["chainId"].(uint64); ok {
		chainID = uint64(chainIDValue)
	}
	
	values := map[string]any{
		"intent_id": intentId,
		"data":      string(message),
	}
	
	if chainID > 0 {
		values["chain_id"] = chainID
	}
	
	values["timestamp"] = time.Now().Unix()
	
	// Use StreamsRedis client for stream operations
	return redis.StreamsRedis.XAdd(ctx, &redisv8.XAddArgs{
		Stream: IntentStream,
		ID:     "*",
		Values: values,
	}).Err()
}
