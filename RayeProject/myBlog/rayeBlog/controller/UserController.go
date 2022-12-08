package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rayeBlog/common"
	"rayeBlog/model"
	"rayeBlog/response"
	"strconv"
)

// Register 注册
func Register(c *gin.Context) {
	db := common.GetDB()
	//获取参数
	var requestUser model.User
	c.Bind(&requestUser)
	userName := requestUser.UserName
	phoneNumber := requestUser.PhoneNumber
	password := requestUser.Password
	//数据验证
	var user model.User
	if userName == "" || phoneNumber == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "信息不能为空",
		})
	} else {
		db.Where("phone_number=?", phoneNumber).First(&user)
		if user.ID != 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": 422,
				"msg":  "用户已存在",
			})
			return
		}
		//密码加密
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		//创建用户
		newUser := model.User{
			UserName:    userName,
			PhoneNumber: phoneNumber,
			Password:    string(hashedPassword),
			Avatar:      "/images/default_avatar.png",
			Collects:    model.Array{},
			Following:   model.Array{},
			Fans:        0,
		}
		db.Create(&newUser)
		//返回结果
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "注册成功",
		})
	}
}

// Login 登录
func Login(c *gin.Context) {
	db := common.GetDB()
	//获取参数
	var requestUser model.User
	c.Bind(&requestUser)
	phoneNumber := requestUser.PhoneNumber
	password := requestUser.Password
	//数据验证
	var user model.User
	db.Where("phone_number=?", phoneNumber).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "用户不存在",
		})
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "密码错误",
		})
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		return
	}
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

// GetInfo 登录后获取信息
func GetInfo(c *gin.Context) {
	//获取上下文中的用户信息
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录获取信息成功",
		"data": gin.H{
			"id":     user.(model.User).ID,
			"avatar": user.(model.User).Avatar,
		},
	})
}

// GetBriefInfo 获取简要信息
func GetBriefInfo(c *gin.Context) {
	db := common.GetDB()
	userId := c.Params.ByName("id")
	user, _ := c.Get("user")
	var curUser model.User
	if userId == strconv.Itoa(int(user.(model.User).ID)) {
		curUser = user.(model.User)
	} else {
		db.Where("id=?", userId).First(&curUser)
		if curUser.ID == 0 {
			response.Fail(c, "用户不存在", nil)
			return
		}
	}
	//返回用户简要信息
	response.Success(c, "查找成功", gin.H{
		"id":      curUser.ID,
		"name":    curUser.UserName,
		"avatar":  curUser.Avatar,
		"loginId": user.(model.User).ID,
	})
}

// GetDetailInfo 获取详细信息
func GetDetailInfo(c *gin.Context) {
	db := common.GetDB()
	userId := c.Params.ByName("id")
	user, _ := c.Get("user")
	var curUser model.User
	if userId == strconv.Itoa(int(user.(model.User).ID)) {
		curUser = user.(model.User)
	} else {
		db.Where("id=?", userId).First(&curUser)
		if curUser.ID == 0 {
			response.Fail(c, "用户不存在", nil)
			return
		}
	}
	//返回用户详细信息
	var articles, collects []model.ArticleInfo
	var following []model.UserInfo
	var collist, follist []string
	collist = ToStringArray(curUser.Collects)
	follist = ToStringArray(curUser.Following)
	db.Table("articles").Select("id, category_id, title, LEFT(content,80) AS content, head_image, "+
		"created_at").Where("user_id = ?", userId).Order("created_at desc").Find(&articles)
	db.Table("articles").Select("id, category_id, title, LEFT(content,80) AS content, head_image, "+
		"created_at").Where("id IN (?)", collist).Order("created_at desc").Find(&collects)
	db.Table("users").Select("id, avatar, user_name").Where("id IN (?)", follist).Find(&following)
	response.Success(c, "查找成功", gin.H{
		"id":        curUser.ID,
		"name":      curUser.UserName,
		"avatar":    curUser.Avatar,
		"loginId":   user.(model.User).ID,
		"articles":  articles,
		"collects":  collects,
		"following": following,
		"fans":      curUser.Fans})

}

// ToStringArray 将自定义类型转换成字符串数组
func ToStringArray(l []string) (a model.Array) {
	for i := 0; i < len(a); i++ {
		l = append(l, a[i])
	}
	return l
}

// ModifyAvatar 修改头像
func ModifyAvatar(c *gin.Context) {
	db := common.GetDB()
	user, _ := c.Get("user")
	var requestUser model.User
	c.Bind(&requestUser)
	avatar := requestUser.Avatar
	//查找用户
	var curUser model.User
	db.Where("id=?", user.(model.User).ID).First(&curUser)
	//更新信息
	if err := db.Model(&curUser).Update("avatar", avatar).Error; err != nil {
		response.Fail(c, "更新失败", nil)
		return
	}
	response.Success(c, "更新成功", nil)
}

// ModifyName 修改用户名
func ModifyName(c *gin.Context) {
	db := common.GetDB()
	// 获取用户ID
	user, _ := c.Get("user")
	// 获取参数
	var requestUser model.User
	c.Bind(&requestUser)
	userName := requestUser.UserName
	// 查找用户
	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)
	// 更新信息
	if err := db.Model(&curUser).Update("user_name", userName).Error; err != nil {
		response.Fail(c, "更新失败", nil)
		return
	}
	response.Success(c, "更新成功", nil)
}
