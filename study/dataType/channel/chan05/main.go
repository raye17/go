package main

import "fmt"

// channel 内置函数
func main() {
	chan01 := make(chan int, 4)
	for i := 0; i < 3; i++ {
		chan01 <- i
	}
	fmt.Println(cap(chan01), len(chan01))
}
