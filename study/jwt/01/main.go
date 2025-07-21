package main

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func main() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NDU4OCwiYWNjb3VudCI6IlY5NDdCaVl6Vk13d1pXOGZ1dXBabnZxQTFtY2lDWUNuayIsImRvbWFpbiI6ImZvbnRyZWUiLCJuaWNrTmFtZSI6IuWwj3RvbnkiLCJwaG9uZSI6IjE3MzExMTExMTE4IiwiZXhwIjoxNzUyNzYwMzYwLCJpc3MiOiJtYWxsIn0.T-WTIa3lrEymFrte1zUwYkoWipOZDN-HloKOfsgjMgg"
	VerifyJWT(tokenStr, secret)
}

var secret = []byte("asdfqwer1234")

func VerifyJWT(tokenStr string, secret []byte) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// 保证算法是我们期望的
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		fmt.Println("❌ 验证失败：", err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("✅ Token 合法，内容如下：")
		for k, v := range claims {
			fmt.Printf("  %s: %v\n", k, v)
		}
	} else {
		fmt.Println("⚠️ Token 无效")
	}
}
