package main

import (
	"fmt"
	"strings"
)

// 匿名函数 闭包
func sai() func() {
	name := "raye"
	return func() {
		fmt.Println("hello,this is func a\n", name)
	}
}
func makeSuffix(prefix string) func(string) string {
	return func(name string) string {
		if !strings.HasPrefix(name, prefix) {
			return prefix + name
		}
		return name
	}
}
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func main() {
	func() {
		fmt.Println("这是匿名函数")
	}()
	s := sai()
	s()
	m := makeSuffix("姓名：")
	fmt.Println(m("raye"))
	add, sub := calc(3)
	fmt.Println(add(1), sub(2))
}
