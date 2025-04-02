package router

import (
	"study/gin/handler"
	"study/gin/middleware"

	"github.com/gin-gonic/gin"
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
