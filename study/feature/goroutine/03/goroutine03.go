package main

// runtime.GOMAXPROCS 限制CPU数量
import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("a", i)
	}
}
func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("b", i)
	}
}
func main() {
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}
