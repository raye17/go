package main

import (
	"fmt"
	"raye/demo/config"
	"raye/demo/db"
	"raye/demo/db/model"
	"raye/demo/pkg/service"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config.AppConfig)
	err = db.InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := service.CreateUser(&model.User{
		Name:   "sss",
		Age:    25,
		Gender: "female",
	}); err != nil {
		fmt.Println(err)
		return
	}
}
