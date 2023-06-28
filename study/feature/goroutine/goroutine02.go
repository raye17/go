package main

import "fmt"

// 并发实现斐波那契数列
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	defer close(c)
}
func main() {
	var chan01 = make(chan int)
	go fibonacci(15, chan01)
	for v := range chan01 {
		fmt.Print(v, " ")
	}
	//messages := make(chan string)
	//go func() {
	//	messages <- "ping"
	//}()
	//msg := <-messages
	//fmt.Println(msg)
}
