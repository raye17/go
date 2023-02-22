package main

import (
	"fmt"
	"time"
)

func write(ch chan string) {
	for i := 0; ; {
		select {
		//写数据
		case ch <- "hello":
			fmt.Println("write hello", i)
			i++
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	output1 := make(chan string, 10)
	go write(output1)
	for s := range output1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}
