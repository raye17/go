package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
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
	dir := filepath.Join("./test0002", "test0001")
	fmt.Println(dir)
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		fmt.Println("mkdir failed", err)
	}

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
