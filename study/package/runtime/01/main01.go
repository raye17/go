package main

//runtime.Gosched() 让出当前时间片，等待重新分配时间片
import (
	"fmt"
	"runtime"
)

func main() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	for i := 0; i < 2; i++ {
		fmt.Println("num:", runtime.NumGoroutine())
		runtime.Gosched()
		fmt.Println("hello")
	}
}
