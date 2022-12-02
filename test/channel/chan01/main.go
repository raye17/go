package main

import "fmt"

// 无缓冲通道（同步通道）
func rec(c chan int) {
	ret := <-c
	fmt.Println("success\n", ret)
}
func main() {
	var ch01 = make(chan int)
	go rec(ch01)
	ch01 <- 10
	//有缓冲通道
	var ch02 = make(chan int, 2)
	ch02 <- 99
	x, ok := <-ch02
	if !ok {
		fmt.Println("closed")
	}
	fmt.Println(x)
}
