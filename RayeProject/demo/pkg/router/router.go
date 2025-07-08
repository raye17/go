package router

import (
	"sxy/demo/pkg/middleware"
	"sxy/demo/pkg/service"
	zaplog "sxy/demo/pkg/zap"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	v1 := r.Group("/v1")
	{
		v1.GET("test", func(c *gin.Context) {
			zaplog.Logger.Info("test")
			c.JSON(200, gin.H{
				"message": "test",
			})
		})
	}
	user := v1.Group("/user")
	{
		user.POST("/register", service.Register)
		user.POST("/login", service.Login)
		user.Use(middleware.CheckLogin())
		user.GET("/info", service.GetUserInfo)
		user.POST("/logout", service.Logout)
	}
	return r
}
