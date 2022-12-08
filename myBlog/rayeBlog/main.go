package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"rayeBlog/common"
	"rayeBlog/routers"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	//配置静态文件路径
	r.StaticFS("/images", http.Dir("static/images"))
	routers.CollectRoutes(r)
	panic(r.Run(":8001"))
}
