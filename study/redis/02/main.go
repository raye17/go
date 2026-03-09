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
	val, err := r.LRange(context.Background(), keys[0], 0, 2).Result()
	if err != nil {
		panic(err)
		return
	}
	for _, key := range keys {
		typeStr, err := r.Type(context.TODO(), key).Result()
		if err != nil {
			fmt.Println("获取类型错误:", err)
			continue
		}
		switch typeStr {
		case "string":
			v, err := r.Get(context.TODO(), key).Result()
			if err != nil {
				fmt.Println("string类型获取失败:", err)
			} else {
				fmt.Println("[string] k:", key, "v:", v)
			}
		case "zset":
			zvals, err := r.ZRangeWithScores(context.TODO(), key, 0, -1).Result()
			if err != nil {
				fmt.Println("zset类型获取失败:", err)
			} else {
				fmt.Printf("[zset] k: %s v: %v\n", key, zvals)
			}
		case "hash":
			hvals, err := r.HGetAll(context.TODO(), key).Result()
			if err != nil {
				fmt.Println("hash类型获取失败:", err)
			} else {
				fmt.Printf("[hash] k: %s v: %v\n", key, hvals)
			}
		default:
			fmt.Println("未处理类型:", typeStr, "key:", key)
		}
	}

	fmt.Println(val)
}
