package main

import (
	"context"
	"fmt"
	"study/redis/cache"

	"github.com/redis/go-redis/v9"
)

var r *redis.Client

func main() {
	r = cache.NewClientRedis()
	keys, err := r.Keys(context.Background(), "*").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)
	vType, _ := r.Type(context.Background(), keys[0]).Result()
	fmt.Println(vType)
	val, err := r.LRange(context.Background(), keys[0], 0, 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
