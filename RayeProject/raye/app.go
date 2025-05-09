package main

import (
	"fmt"
	"raye/demo/config"
	"raye/demo/db"
	"raye/demo/pkg/aliyun"
	"raye/demo/pkg/mq"
	"raye/demo/pkg/router"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		fmt.Println("load config failed, err: ", err)
		return
	}
	fmt.Println(config.AppConfig.System.Mode, config.AppConfig.System.Version)
	err = db.InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	// redis := cache.RedisClient
	// s, err := redis.Set(context.Background(), "sss", "my name is lalla", time.Hour*24).Result()
	// if err != nil {
	// 	fmt.Println("key sss set failed, err:", err)
	// 	return
	// }
	//fmt.Println("s: ", s)
	aliyun.NewOss()
	go mq.Listening()
	router.NewRouter()
}
