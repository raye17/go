package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"loginDemo/common"
	"loginDemo/model"
	"net/http"
)

func Register(ctx *gin.Context) {
	db := common.GetDb()
	//获取参数
	//此处使用Bind()函数，可以处理不同格式的前端数据
	var requestUser model.User
	err := ctx.Bind(&requestUser)
	if err != nil {
		return
	}
	name := requestUser.Name
	password := requestUser.Password
	//数据验证
	if len(name) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名不能为空",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于六位",
		})
		return
	}
	var user model.User
	db.Where("name=?", name).First(&user)
	if user.ID != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    "422",
			"message": "用户已存在",
		})
		return
	}
	//创建用户
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		return
	}
	newUser := model.User{
		Name:     name,
		Password: string(hashPassword),
	}
	db.Create(&newUser)
	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})
}

// Login 登录
func Login(ctx *gin.Context) {
	db := common.GetDb()

	//获取参数
	//此处使用Bind()函数，可以处理不同格式的前端数据
	var requestUser model.User
	err := ctx.Bind(&requestUser)
	if err != nil {
		return
	}
	name := requestUser.Name
	password := requestUser.Password
	//数据验证
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}
	var user model.User
	db.Where("name=?", name).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    "422",
			"message": "用户不存在",
		})
		return
	}
	//判断密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码错误",
		})

	} else {
		//发放token
		token, err := common.ReleaseToken(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "系统异常",
			})
			//记录下错误
			log.Printf("token generate error: %v", err)
			return
		} else {
			//返回结果
			ctx.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": gin.H{
					"token": token,
				},
				"message": "登录成功!",
			})
		}
	}
}
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	//将用户信息返回
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user": user,
		},
	})
}
