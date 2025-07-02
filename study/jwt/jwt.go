package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	secretKey = "hello world"
	// 指定token过期时间
	tokenExpireDuration = time.Hour * 24
	// 指定refresh token过期时间
	refreshTokenExpireDuration = time.Hour * 24 * 7 * 30
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个UserID字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT
func GenerateToken(userID int64) (string, string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireDuration).Unix(), // 过期时间
			Issuer:    "bluebell",                                 // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", "", err
	}

	// 生成刷新令牌
	rc := MyClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTokenExpireDuration).Unix(),
			Issuer:    "bluebell",
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rc)
	newRefreshToken, err := refreshToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新AccessToken
func RefreshToken(accessToken, refreshToken string) (string, string, error) {
	// 1. 解析refreshToken，如果无效则返回错误
	claims, err := ParseToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	// 2. 检查refreshToken是否已过期
	if time.Now().Unix() > claims.ExpiresAt {
		return "", "", errors.New("refresh token has expired")
	}

	// 3. 生成新的accessToken和refreshToken
	return GenerateToken(claims.UserID)
}
