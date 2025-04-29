package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

// GenerateToken 签发用户Token
func GenerateToken(id uint, name string, hour int, jwtSecret []byte) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(hour) * time.Hour)
	claims := Claims{
		ID:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "mall",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
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
