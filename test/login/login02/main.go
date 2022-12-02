package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"loginDemo/common"
	"loginDemo/route"
)

func main() {
	//获取初始化数据库
	db := common.InitDb()
	defer db.Close()
	r := gin.Default()
	route.CollectRoutes(r)
	panic(r.Run(":9090"))
}
