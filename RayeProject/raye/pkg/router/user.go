package router

import (
	"raye/demo/pkg/service"

	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	//noAuth := r.Group("")
	auth := r.Group("")
	user := auth.Group("/user")
	{
		user.POST("create", service.CreateUser)
		user.POST("detail", service.GetUserByID)
		user.POST("list", service.GetUserList)
		user.POST("update", service.UpdateUser)
		user.DELETE("delete", service.DeleteUser)
	}
}
