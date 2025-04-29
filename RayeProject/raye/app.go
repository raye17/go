package main

import (
	"encoding/json"
	"fmt"
	"raye/demo/config"
	"raye/demo/db"
	"raye/demo/pkg/mq"
	"raye/demo/pkg/router"
)

type userInfo struct {
	Id        uint
	Name      string
	Age       int
	Gender    string
	CreatedAt string
	UpdatedAt string
}

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
	go mq.Listening()
	router.NewRouter()
}
func (u userInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
