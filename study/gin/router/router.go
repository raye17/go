package router

import (
	"github.com/gin-gonic/gin"
	"sxy/demo/gin/handler"
	"sxy/demo/gin/middleware"
)

func NewRouter() *gin.Engine {

	// 注册路由
	r := gin.Default()

	r.POST("/login", handler.LoginHandler)
	auth := r.Group("/api")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/user/info", handler.GetInfoFromC)
	}
	return r

}
