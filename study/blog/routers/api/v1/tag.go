package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"rayeBlog/models"
	"rayeBlog/pkg/setting"
	"rayeBlog/pkg/types"
	"rayeBlog/pkg/util"
)

// GetTags 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := types.SUCCESS
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  types.GetMsg(code),
		"data": data,
	})
}

// AddTags 新增文章标签
func AddTags(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许为0或1")
	code := types.InvalidParams
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = types.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = types.ErrorExistTag
		}
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  types.GetMsg(code),
		"data": make(map[string]string),
	})
}

// EditTags 修改文章标签
func EditTags(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")
	valid := validation.Validation{}
	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许为0或1")
		valid.Required(id, "id").Message("ID不能为空")
		valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
		valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
		valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	}
	code := types.InvalidParams
	if !valid.HasErrors() {
		code = types.SUCCESS
		if models.ExistTagById(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			models.EditTag(id, data)
		} else {
			code = types.ErrorNotExistTag
		}
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  types.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteTags 删除文章标签
func DeleteTags(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id 必须大于0")
	code := types.InvalidParams
	if !valid.HasErrors() {
		code = types.SUCCESS
		if models.ExistTagById(id) {
			models.DeleteTag(id)
		} else {
			code = types.ErrorNotExistTag
		}
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  types.GetMsg(code),
		"data": make(map[string]string),
	})
}
