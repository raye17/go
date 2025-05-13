package main

import (
	"fmt"
	"study/redis/cache"
)

func main() {
	redis := cache.NewClientRedis()
	fmt.Println(redis)
}
