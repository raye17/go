package routes

import (
	"github.com/gin-gonic/gin"
	"study/package/github.com/gin/ginGorm/controller"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	/*
		用户User路由组
	*/
	userGroup := r.Group("user")
	{
		//增加
		userGroup.POST("/users", controller.CreateUser)
		//查看所有User
		userGroup.GET("/users", controller.GetUserList)
		//修改
		userGroup.PUT("/users/:id", controller.UpdateUser)
		//删除
		userGroup.DELETE("/users/:id", controller.DeleteUserById)
	}
	return r
}
