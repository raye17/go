package common

import (
	"github.com/dgrijalva/jwt-go"
	"rayeBlog/model"
	"time"
)

// jwt加密密钥
var jwtKey = []byte("raye")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// ReleaseToken 生成token
func ReleaseToken(user model.User) (string, error) {
	//token有效期
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			//过期时间
			ExpiresAt: expirationTime.Unix(),
			//发放时间
			IssuedAt: time.Now().Unix(),
			Issuer:   "raye",
		},
	}
	//使用jwt秘钥生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	//返回token
	return tokenString, nil
}
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
