package main

import "fmt"

func f1(ch1 chan int) {
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}
func f2(ch1 chan int, ch2 chan int) {
	for {
		temp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- temp * temp
	}
	close(ch2)
}
func main() {
	var ch01 = make(chan int, 100)
	var ch02 = make(chan int, 100)
	go f1(ch01)
	go f2(ch01, ch02)
	for ret := range ch02 {
		fmt.Println(ret)
	}
}
