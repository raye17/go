package main

import (
	"context"
	"fmt"
	"study/db/redis/client"
	"study/util/errors"
)

// redis set

func main() {
	c := client.NewClientRedis()
	defer c.Close()
	err := c.Set(context.Background(), "abc", 123, 0).Err()
	errors.Checkout("redis set key:'abc' failed!", err)
	res, err := c.Get(context.Background(), "abc").Result()
	errors.Checkout("redis get key:'abc' failed", err)
	fmt.Println("result:", res)
	err = c.Set(context.Background(), "name", "sxy", 0).Err()
	errors.Checkout("redis set key 'name' failed", err)
}
