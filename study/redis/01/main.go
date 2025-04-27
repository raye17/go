package main

import (
	"context"
	"encoding/json"
	"fmt"
	"study/redis/cache"
)

type user struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	u1 := user{
		Name:   "sss",
		Age:    19,
		Gender: "女",
	}
	u2 := user{
		Name:   "002",
		Age:    199,
		Gender: "female",
	}
	c := cache.NewClientRedis()
	jsu1, err := json.Marshal(u1)
	if err != nil {
		panic(err)
	}
	err = c.Set(context.Background(), "user01", jsu1, 0).Err()
	if err != nil {
		panic(err)
	}
	jsu2, err := json.Marshal(u2)
	err = c.Set(context.Background(), "user02", jsu2, 0).Err()
	if err != nil {
		panic(err)
	}
	//val, err := c.Get(context.Background(), "raye").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("raye:", val)
	val, err := c.Get(context.Background(), "user01").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("user01:", val)
	//err = c.Del(context.Background(), "user01").Err()
	//if err != nil {
	//	panic(err)
	//}
	err = c.Close()
	if err != nil {
		panic(err)
	}
}
