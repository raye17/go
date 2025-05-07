package service

import (
	"fmt"
	"ssh/demo/api/dci"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

var GrpcDciImpl = new(dci.DciClientImpl)

func init() {
	fmt.Println("init dubbogo service")
	config.SetConsumerService(GrpcDciImpl) //dci
	if err := config.Load(); err != nil {
		panic(err)
	}
}
func Ser() {
	fmt.Println("init service import")
}
