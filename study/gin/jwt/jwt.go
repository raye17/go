package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

// 定义密钥
var jwtSecret = []byte("sss")

// Claims 自定义JWT的声明结构
type Claims struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT token
func GenerateToken(user string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(24) * time.Hour)
	// 创建自定义声明
	claims := Claims{
		Name:     user,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			// 过期时间设置为24小时
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  nowTime.Unix(),
			NotBefore: nowTime.Unix(),
		},
	}

	// 使用指定的签名方法创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名并获得完整的编码后的字符串token
	return token.SignedString(jwtSecret)
}

// ParseToken 解析JWT token
func ParseToken(tokenString string) (*Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// 验证token是否有效
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ValidateToken 验证token是否有效
func ValidateToken(tokenString string) bool {
	_, err := ParseToken(tokenString)
	return err == nil
}
