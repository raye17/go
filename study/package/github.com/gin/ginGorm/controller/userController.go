package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study/package/github.com/gin/ginGorm/entity"
	"study/package/github.com/gin/ginGorm/service"
)

// CreateUser 创建User
func CreateUser(c *gin.Context) {
	//定义一个User变量
	var user entity.User
	//将调用后端的request请求中body数据根据json格式解析到User结构变量中
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	//将被转换的user变量传给service层的CreateUser方法，进行User的新建
	err = service.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": user,
		})
	}
}

// GetUserList 查询User
func GetUserList(c *gin.Context) {
	todoList, err := service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": todoList,
		})
	}
}

// UpdateUser 修改User
func UpdateUser(c *gin.Context) {

}

// DeleteUserById 删除User
func DeleteUserById(c *gin.Context) {

}
