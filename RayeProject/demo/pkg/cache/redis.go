package cache

import (
	"sxy/demo/config"
	zaplog "sxy/demo/pkg/zap"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.Redis.Addr,
		Password: config.AppConfig.Redis.Password,
		DB:       config.AppConfig.Redis.DB,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		zaplog.Logger.Error("connect redis failed, err:", zap.Error(err))
		return
	}
	zaplog.Logger.Info("connect redis success")
}
