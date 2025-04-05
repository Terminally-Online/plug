package redis

import (
	"fmt"
	"solver/internal/utils"
)

var (
	RedisHost    = utils.GetEnvOrDefault("REDIS_HOST", "localhost")
	RedisPort    = utils.GetEnvOrDefault("REDIS_POST", "6379")
	RedisAddress = fmt.Sprintf("%s:%s", RedisHost, RedisPort)
)

