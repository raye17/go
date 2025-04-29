package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"raye/demo/db"
	"raye/demo/db/model"
	"raye/demo/pkg/cache"
	"raye/demo/pkg/service"
	"raye/demo/pkg/utlis/e"
	"raye/demo/pkg/utlis/jwt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const Authorization = "Authorization"

// JWTMiddleware JWT token验证中间件
func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(Authorization)
		fmt.Println(authHeader)
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is empty"})
			return
		}

		// // 确保Authorization header格式正确并提取token
		// splitToken := strings.Split(authHeader, "Bearer ")
		// if len(splitToken) != 2 {
		// 	fmt.Println("22222222")
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		// 	return
		// }

		// token := splitToken[1]
		username, err := verifyToken(authHeader)
		if err != nil {
			fmt.Println("33333333", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		var user model.User
		if err := db.DbTest01.Table("user").Where("name =?", username).First(&user).Error; err != nil {
			if err.Error() == gorm.ErrRecordNotFound.Error() {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			return
		}
		userInfo := service.UserInfo{
			Id:     user.ID,
			Age:    user.Age,
			Gender: user.Gender,
			Token:  authHeader,
		}

		// 将认证的用户名添加到请求上下文
		ctx.Set("userInfo", userInfo)
		ctx.Next()
	}
}

// verifyToken 验证JWT token并返回用户名
func verifyToken(tokenString string) (string, error) {
	// 检查token是否在黑名单中
	isBlacklisted, err := cache.RedisClient.SIsMember(context.Background(), "jwt:blacklist", tokenString).Result()
	if err != nil {
		return "", err
	}
	if isBlacklisted {
		return "", errors.New("token is blacklisted")
	}

	token, err := jwt.ParseToken(tokenString, e.JWTSecret)
	if err != nil {
		return "", err
	}

	return token.Name, nil
}
