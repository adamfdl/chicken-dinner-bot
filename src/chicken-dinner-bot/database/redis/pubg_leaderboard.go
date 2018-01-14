package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type PUBGLeaderboardOperator interface {
	RetrieveLeaderBoard() ([]redis.Z, error)
	AddNewPlayer(string, string) error
	IncrementPlayerScore(string, string, int)
}

var pubgLeaderboardOperator PUBGLeaderboardOperator

type PUBGLeaderboardOperatorImpl struct{}

func GetPUBGLeaderboardOperator() PUBGLeaderboardOperator {
	if pubgLeaderboardOperator == nil {
		return &PUBGLeaderboardOperatorImpl{}
	}
	return pubgLeaderboardOperator
}

func SetPUBGLeaderboardOperator(instance PUBGLeaderboardOperator) {
	pubgLeaderboardOperator = instance
}

//
func (*PUBGLeaderboardOperatorImpl) RetrieveLeaderBoard() ([]redis.Z, error) {
	if err := shouldPerformOnRedis(); err != nil {
		logrus.Warn("[REDIS]", err.Error())
		return nil, err
	}

	key := os.Getenv("REDIS_SORTED_SET_KEY")
	if result, err := redisClient.ZRevRangeByScoreWithScores(key, redis.ZRangeBy{Max: "+inf", Min: "-inf"}).Result(); err == nil {
		return result, nil
	} else {
		logrus.Warn("[REDIS] Cannot retrieve", key, "Error: ", err.Error())
		return nil, err
	}
}

func (*PUBGLeaderboardOperatorImpl) AddNewPlayer(discord_id, pubg_nick string) error {
	if err := shouldPerformOnRedis(); err != nil {
		logrus.Warnf("[REDIS] %s", err.Error())
		return err
	}

	member := fmt.Sprintf("%s:%s", discord_id, pubg_nick)
	if err := redisClient.ZAdd("pubg_leaderboard", redis.Z{Member: member, Score: 0}); err != nil {
		logrus.Warnf("[REDIS] Cannot store new player from redis. Error: %s", err.Err())
		return err.Err()
	}

	return nil
}

func (*PUBGLeaderboardOperatorImpl) IncrementPlayerScore(discord_id, pubg_nick string, score int) {
	if err := shouldPerformOnRedis(); err != nil {
		logrus.Warnf("[REDIS] %s", err.Error())
		return
	}

	key := fmt.Sprintf("discordID:%s-pubgNick:%s", discord_id, pubg_nick)
	if err := redisClient.ZIncrBy(key, float64(score), "pubg_leaderboard"); err != nil {
		logrus.Warnf("[REDIS] Cannot increment this member. Error: %s", err.Err())
	}
}
