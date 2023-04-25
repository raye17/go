package main

import (
	"fmt"
	"study/util/encode"
	"study/util/encrypt"
)

func main() {
	p := "ODg4ODg4"
	s := encode.CodeDecode(p)
	pp, _ := encrypt.Encrypt(s)
	fmt.Println(s)
	fmt.Println(string(pp))
}
