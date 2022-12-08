package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rayeBlog/common"
	"rayeBlog/model"
	"strings"
)

/*
编写一个中间件，获取到前端请求中的token，调用ParseToken()对其进行解析，
若token不合规范，该请求将会被抛弃，当token符合规范时才可以进行下一步操作
*/

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取authorization header
		tokenString := c.Request.Header.Get("Authorization")
		//token为空
		if tokenString == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		//非法token
		if len(tokenString) < 7 || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		//提取token的有效部分
		tokenString = tokenString[7:]
		//解析token
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		//获取claims中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.Where("id=?", userId).First(&user)
		//将用户信息写入上下文便于读取
		c.Set("user", user)
		c.Next()
	}
}
