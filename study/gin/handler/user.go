package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"sxy/demo/gin/jwt"
)

// 用于存储验证码的map（实际项目中应该使用Redis）
var verificationCodes = make(map[string]string)

// LoginRequest 登录请求参数
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SendCodeRequest 发送验证码请求
type SendCodeRequest struct {
	Phone string `json:"phone" binding:"required,len=11"`
}

// LoginByCodeRequest 验证码登录请求
type LoginByCodeRequest struct {
	Phone string `json:"phone" binding:"required,len=11"`
	Code  string `json:"code" binding:"required,len=6"`
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"required"`
}

// LoginHandler 处理登录请求
func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 验证用户名和密码
	// 用户名为admin，密码为123456
	if req.Username == "admin" && req.Password == "123456" {
		// 生成Token
		token, err := jwt.GenerateToken(req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "生成token失败",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "登录成功",
			"data": gin.H{
				"token": token,
			},
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"code": 401,
		"msg":  "用户名或密码错误",
	})
}

// GetInfoFromC 获取用户信息
func GetInfoFromC(c *gin.Context) {
	// 从上下文中获取用户信息

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{},
	})
}

// SendVerificationCode 发送验证码
func SendVerificationCode(c *gin.Context) {
	var req SendCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 生成6位随机验证码
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	// 将验证码保存到map中（实际项目中应该使用Redis，并设置过期时间）
	verificationCodes[req.Phone] = code

	// 应该调用短信服务发送验证码
	// 这里直接返回验证码
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "发送成功",
		"data": gin.H{
			"code": code,
		},
	})
}

// LoginByCode 验证码登录
func LoginByCode(c *gin.Context) {
	var req LoginByCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 验证验证码
	savedCode, exists := verificationCodes[req.Phone]
	if !exists || savedCode != req.Code {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "验证码错误或已过期",
		})
		return
	}

	// 验证通过后删除验证码
	delete(verificationCodes, req.Phone)

	// 生成Token
	token, err := jwt.GenerateToken(req.Phone, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}
