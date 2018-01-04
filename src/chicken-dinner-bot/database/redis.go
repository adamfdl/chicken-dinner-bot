package database

import (
	"os"

	"github.com/go-redis/redis"
)

func GetRedisClient() (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_SERVER_IP"),
		Password: os.Getenv("REDIS_SERVER_PASSWORD"),
		DB:       0,
	})

	if _, err := redisClient.Ping().Result(); err != nil {
		return nil, err
	}

	return redisClient, nil
}
