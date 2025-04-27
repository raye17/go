package router

import "github.com/gin-gonic/gin"

func NewRouter() {
	r := gin.Default()

	// 注册路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	privateGroup := r.Group("")
	userRouter(privateGroup)
	// 启动服务器
	r.Run(":8080")
}
