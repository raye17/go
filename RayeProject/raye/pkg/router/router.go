package router

import (
	"raye/demo/pkg/middleware"
	"raye/demo/pkg/service"

	"github.com/gin-gonic/gin"
)

func NewRouter() {
	r := gin.Default()
	r.Use(middleware.Cors())
	// 注册路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r.POST("user/login", service.Login)
	r.Use(middleware.JWTMiddleware())
	r.POST("user/logout", service.Logout)
	privateGroup := r.Group("")
	userRouter(privateGroup)
	// 启动服务器
	r.Run(":8080")
}
