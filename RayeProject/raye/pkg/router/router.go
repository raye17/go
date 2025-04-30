package router

import (
	"net/http"
	"raye/demo/pkg/middleware"

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
	privateGroup := r.Group("")
	privateGroup.Use(middleware.JWTMiddleware())
	userRouter(privateGroup)
	ImgRouter(privateGroup)
	r.StaticFS("/static", http.Dir("./runtime"))
	r.StaticFile("/favicon.ico", "./runtime/favicon.ico")
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 1,
			"msg":    "不存在的路由",
		})
	})
	// 启动服务器
	r.Run(":8080")
}
