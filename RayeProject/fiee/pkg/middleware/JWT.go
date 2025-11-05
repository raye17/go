package middleware

import (
	"test001/pkg/e"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := ParseToken(token, e.JWTSecret)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"code":   code,
				"status": e.InvalidToken,
				"msg":    e.GetMsg(e.InvalidToken),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

type Claims struct {
	ID        uint   `json:"id"`
	Account   string `json:"account"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

// ParseToken 验证用户token
func ParseToken(token string, jwtSecret []byte) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
