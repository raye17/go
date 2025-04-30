package router

import (
	"raye/demo/pkg/middleware"
	"raye/demo/pkg/service"

	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	//noAuth := r.Group("")
	r.POST("user/login", service.Login)

	auth := r.Group("")
	auth.Use(middleware.JWTMiddleware())
	user := auth.Group("/user")
	{
		user.POST("create", service.CreateUser)
		user.POST("detail", service.GetUserByID)
		user.GET("list", service.GetUserList)
		user.POST("update", service.UpdateUser)
		user.DELETE("delete", service.DeleteUser)
		auth.POST("user/logout", service.Logout)
	}
}
