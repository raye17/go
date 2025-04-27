package main

import (
	"fmt"
	"raye/demo/config"
	"raye/demo/db"
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
}
