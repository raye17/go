package cache

import (
	"context"

	"github.com/dubbogo/gost/log/logger"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func init() {
	clientRedis := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	_, err := clientRedis.Ping(context.Background()).Result()
	if err != nil {
		logger.Errorf("connRedis err :%v", err)
		panic(err)
	}
	RedisClient = clientRedis
}
func NewClientRedis() *redis.Client {
	return RedisClient
}
