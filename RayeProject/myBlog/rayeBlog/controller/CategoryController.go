package controller

import (
	"github.com/gin-gonic/gin"
	"rayeBlog/common"
	"rayeBlog/model"
	"rayeBlog/response"
)

// SearchCategory 查询分类
func SearchCategory(c *gin.Context) {
	db := common.GetDB()
	var categories []model.Category
	if err := db.Find(&categories).Error; err != nil {
		response.Fail(c, "查找失败", nil)
		return
	}
	response.Success(c, "查找成功", gin.H{
		"categories": categories,
	})

}

// SearchCategoryName 查询分类名
func SearchCategoryName(c *gin.Context) {
	db := common.GetDB()
	var category model.Category
	//获取path中的分类id
	categoryId := c.Params.ByName("id")
	if err := db.Where("id=?", categoryId).First(&category).Error; err != nil {
		response.Fail(c, "分类不存在", nil)
		return
	}
	response.Success(c, "查找成功", gin.H{
		"categoryName": category.CategoryName,
	})
}
