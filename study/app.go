package main

import (
	"context"
	"fmt"
	"study/redis/cache"
)

func main() {
	redis := cache.NewClientRedis()
	fmt.Println(redis)
	keys, err := redis.Keys(context.TODO(), "*").Result()
	if err != nil {
		fmt.Println(err)
	}
	for _, key := range keys {
		typeStr, err := redis.Type(context.TODO(), key).Result()
		if err != nil {
			fmt.Println("获取类型错误:", err)
			continue
		}
		switch typeStr {
		case "string":
			v, err := redis.Get(context.TODO(), key).Result()
			if err != nil {
				fmt.Println("string类型获取失败:", err)
			} else {
				fmt.Println("[string] k:", key, "v:", v)
			}
		case "zset":
			zvals, err := redis.ZRangeWithScores(context.TODO(), key, 0, -1).Result()
			if err != nil {
				fmt.Println("zset类型获取失败:", err)
			} else {
				fmt.Printf("[zset] k: %s v: %v\n", key, zvals)
			}
		case "hash":
			hvals, err := redis.HGetAll(context.TODO(), key).Result()
			if err != nil {
				fmt.Println("hash类型获取失败:", err)
			} else {
				fmt.Printf("[hash] k: %s v: %v\n", key, hvals)
			}
		default:
			fmt.Println("未处理类型:", typeStr, "key:", key)
		}
	}
}
