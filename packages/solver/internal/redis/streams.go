package redis

import (
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	StreamDatabase    = 1
	StreamDialTimeout = 1 * time.Second
	StreamReadTimeout = 30 * time.Second
	StreamsRedis      = redis.NewClient(&redis.Options{
		Addr:        RedisAddress,
		Password:    os.Getenv("REDIS_PASSWORD"),
		DB:          StreamDatabase,
		DialTimeout: StreamDialTimeout,
		ReadTimeout: StreamReadTimeout,
	})
)
