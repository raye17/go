package route

import (
	"github.com/gin-gonic/gin"
	"loginDemo/controller"
	"loginDemo/middleware"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
