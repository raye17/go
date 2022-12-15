package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"loginDemo/common"
	"loginDemo/route"
)

func main() {
	//获取初始化数据库
	err := common.InitDb()
	if err != nil {
		fmt.Println(err)
	}
	r := gin.Default()
	route.CollectRoutes(r)
	panic(r.Run(":9090"))
}
