package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type user struct {
	name   string
	age    int
	gender string
}

func main() {
	//user := &user{
	//	name:   "",
	//	age:    19,
	//	gender: "man",
	//}
	////a := os.FileMode(0777).String()
	//switch {
	//case user.name == "raye":
	//	fmt.Println("name")
	//case user.age == 0:
	//	fmt.Println("age")
	//case user.gender == "":
	//	fmt.Println("gender")
	//}
	//fmt.Println("over")
	//s1 := []string{"raye.lello"}
	//s2 := []string{"raye,hello"}
	//l := CompareSlice(s1, s2)
	//fmt.Println(l)
	//dir := filepath.Join("./test0002", "test0001")
	//fmt.Println(dir)
	//err := os.MkdirAll(dir, 0777)
	//if err != nil {
	//	fmt.Println("mkdir failed", err)
	//}
	pwd := []byte("123445656")
	user := "sss"
	hashPwd, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(hashPwd)))
	dstuser := make([]byte, base64.StdEncoding.EncodedLen(len([]byte(user))))
	base64.StdEncoding.Encode(dst, hashPwd)
	base64.StdEncoding.Encode(dstuser, []byte(user))
	fmt.Println(string(dstuser))
	fmt.Println(string(dst))
	pp, _ := base64.StdEncoding.DecodeString(string(dst))
	fmt.Println(string(pp))
	//err := bcrypt.CompareHashAndPassword(pp, pwd)
	//if err != nil {
	//	fmt.Println("err")
	//	return
	//} else {
	//	fmt.Println("ok")
	//}
}
func info(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    90,
		"message": "登录成功",
	})
	fmt.Println("info")
}

func CompareSlice(str1, str2 []string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for _, s1 := range str1 {
		found := false
		for _, s2 := range str2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}

	}
	return false
}
