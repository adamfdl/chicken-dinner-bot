package redis

import (
	"chicken-dinner-bot/database"

	"github.com/go-redis/redis"
)

var (
	redisClient          *redis.Client
	redisConnectionError error
)

func shouldPerformOnRedis() error {
	if redisClient != nil {
		_, redisConnectionError = redisClient.Ping().Result()
	}

	if redisClient == nil || redisConnectionError != nil {
		redisClient, redisConnectionError = database.GetRedisClient()
	}

	return redisConnectionError
}
