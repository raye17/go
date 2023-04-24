package main

import "fmt"

// 字符串的修改
// 字符串不能直接修改
func main() {
	s := "raye"
	s1 := []byte(s)
	s1[0] = 'R'
	fmt.Println("byte(s1):", s1, "string(s1):", string(s1))
}
