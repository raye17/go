package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"sxy/demo/pkg/cache"
	"sxy/demo/pkg/common/jwt"
	"sxy/demo/pkg/model/login"

	"github.com/gin-gonic/gin"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		claims, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}
		key := fmt.Sprintf(login.BlackLoginKey, claims.UserId, claims.Username)
		if err := cache.RedisClient.Get(key).Err(); err == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "token已失效",
			})
			c.Abort()
			return
		}
		userInfo := login.UserInfo{
			UserId:   claims.UserId,
			Username: claims.Username,
			Token:    token,
		}
		c.Set("jwtInfo", userInfo)
	}
}
