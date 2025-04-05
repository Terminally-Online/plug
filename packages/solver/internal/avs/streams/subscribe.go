package streams

import (
	"context"
	"solver/internal/redis"
	"time"

	redisv8 "github.com/go-redis/redis/v8"
)

func Subscribe(ctx context.Context, consumerGroup, consumer string) <-chan redisv8.XStream {
	_ = redis.StreamsRedis.XGroupCreateMkStream(ctx, IntentStream, consumerGroup, "0").Err()
	channel := make(chan redisv8.XStream)

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(channel)
				return
			default:
				streams, err := redis.StreamsRedis.XReadGroup(ctx, &redisv8.XReadGroupArgs{
					Group:    consumerGroup,
					Consumer: consumer,
					Streams:  []string{IntentStream, ">"},
					Count:    10,
					Block:    time.Second * 5,
				}).Result()

				if err != nil && err != redisv8.Nil {
					time.Sleep(time.Second)
					continue
				}

				for _, stream := range streams {
					channel <- stream
				}
			}
		}
	}()

	return channel
}
