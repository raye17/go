package cache

import (
	"github.com/go-redis/redis"
	"github.com/google/wire"
)

var RedisProvider = wire.NewSet(NewRedis)

// TODO 添加连接池
func NewRedis() *redis.Client {
	//redisDb, _ := strconv.Atoi(dciConfig.Data.Redis.DB)
	//RedisClient := redis.NewClient(&redis.Options{
	//	Addr:     dciConfig.Data.Redis.Addr,
	//	Password: dciConfig.Data.Redis.Password,
	//	DB:       redisDb,
	//})
	//_, err := RedisClient.Ping().Result()
	//if err != nil {
	//	logger.Errorf("connRedis err", err)
	//	panic(err)
	//}
	return nil
}
