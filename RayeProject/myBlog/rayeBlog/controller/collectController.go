package controller

import (
	"github.com/gin-gonic/gin"
	"rayeBlog/common"
	"rayeBlog/model"
	"rayeBlog/response"
	"strconv"
)

// Collects 查询收藏
func Collects(c *gin.Context) {
	db := common.GetDB()
	user, _ := c.Get("user")
	id := c.Params.ByName("id")
	var curUser model.User
	db.Where("id=?", user.(model.User).ID).First(&curUser)
	//判断是否已收藏
	for i := 0; i < len(curUser.Collects); i++ {
		if curUser.Collects[i] == id {
			response.Success(c, "查询成功", gin.H{
				"collected": true,
				"index":     i,
			})
			return
		}
	}
	response.Success(c, "查询成功", gin.H{
		"collected": false,
	})
}

// NewCollect 新增收藏
func NewCollect(c *gin.Context) {
	db := common.GetDB()
	// 获取用户ID
	user, _ := c.Get("user")
	// 获取path中的id
	id := c.Params.ByName("id")
	// 查找用户
	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)
	var newCollects []string
	newCollects = append(curUser.Collects, id)
	// 更新收藏夹
	if err := db.Model(&curUser).Update("collects", newCollects).Error; err != nil {
		response.Fail(c, "更新失败", nil)
		return
	}
	response.Success(c, "更新成功", nil)
}

// UnCollect 取消收藏
func UnCollect(c *gin.Context) {
	db := common.GetDB()
	// 获取用户ID
	user, _ := c.Get("user")
	// 获取path中的index
	index, _ := strconv.Atoi(c.Params.ByName("index"))
	// 查找用户
	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)
	var newCollects []string
	newCollects = append(curUser.Collects[:index], curUser.Collects[index+1:]...)
	// 更新收藏夹
	if err := db.Model(&curUser).Update("collects", newCollects).Error; err != nil {
		response.Fail(c, "更新失败", nil)
		return
	}
	response.Success(c, "更新成功", nil)
}
