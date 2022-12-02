package handler

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"raye/blog/config"
	"raye/blog/models"
	"raye/blog/types"
	"raye/blog/util"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 7 * 24

var Secret = []byte("gin-api-blog")

// GenToken 生成token
func GenToken(username string) (string, error) {
	claims := Claims{
		"username", // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "raye",                                     // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(Secret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// AuthLogin 登录认证
func AuthLogin(c *gin.Context) {
	// 用户发送用户名和密码过来
	username := c.PostForm("username")
	password, err := util.PasswordHash(c.PostForm("password"))
	if err != nil {
		fmt.Println("err:", err)
	}

	db := config.GetDB()

	type UserInfo struct {
		Username string
		Password string
	}
	userInfo := make([]UserInfo, 2)
	db.First(&models.User{}).Where("username = ? AND password = ?", username, password).Scan(&userInfo)
	fmt.Println(userInfo)
	// 校验用户名和密码是否正确
	if username == userInfo[0].Username && password == userInfo[0].Password {
		// 生成Token
		tokenString, _ := GenToken(username)
		util.Success(c, tokenString)
		return
	}
	util.Error(c, int(types.ApiCode.AUTHFAILED), types.ApiCode.GetMessage(types.ApiCode.AUTHFAILED))
	return
}
