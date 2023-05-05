package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var r *redis.Client

func init() {
	r = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}
func main() {
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
