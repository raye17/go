package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"time"
)

type myClaims struct {
	jwt.StandardClaims
	user   string
	gender string
}

func main() {
	username := "raye"
	expire := time.Now().Add(time.Minute * 30).Unix()
	myClaims := &myClaims{
		jwt.StandardClaims{
			ExpiresAt: expire,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "sxy",
			Subject:   username,
		},
		username,
		"male",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	fmt.Printf("%+v\n", *token)
	fmt.Println(reflect.ValueOf(token))
	signedToken, err := token.SignedString([]byte("sss"))
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("s:", signedToken)
}
