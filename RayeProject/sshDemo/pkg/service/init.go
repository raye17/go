package service

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

func init() {
	//config.SetConsumerService() //dci

	if err := config.Load(); err != nil {
		panic(err)
	}
}
