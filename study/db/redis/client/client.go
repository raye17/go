package client

import "github.com/redis/go-redis/v9"

var clientRedis *redis.Client

func init() {
	clientRedis = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}
func NewClientRedis() *redis.Client {
	return clientRedis
}
