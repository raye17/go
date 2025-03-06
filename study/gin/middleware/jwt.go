package middleware

import (
	"net/http"
	"strings"
	"sxy/demo/gin/jwt"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Authorization 头获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return

		}

		// parts[1]是获取到的tokenString，使用之前定义好的解析JWT的函数来解析它
		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}

		// 将当前请求的用户信息保存到请求的上下文c上
		c.Set("name", claims.Name)
		c.Set("age", claims.Password)
		c.Next()
	}
}
