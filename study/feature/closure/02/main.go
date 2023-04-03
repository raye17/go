package main

import (
	"fmt"
	"time"
)

func main() {
	f := timeSpeed(slowFunc)
	f()
}
func timeSpeed(fn func()) func() {
	fmt.Println("this is timeSpeed...")
	return func() {
		fmt.Println("this is return...")
		start := time.Now()
		fn()
		fmt.Println("after fn()")
		fmt.Println("time speed:", time.Since(start).Seconds())
	}
}
func slowFunc() {
	fmt.Println("this is slow func")
	time.Sleep(time.Second)
	fmt.Println("after 1 second...")
}
