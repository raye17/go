package main

import (
	"fmt"
	"time"
)

func say() {
	time.Sleep(time.Second)
	fmt.Println("say")
	go func() {
		time.Sleep(time.Second)
		fmt.Println("say:")
		fmt.Println("func..")
	}()
}

func main() {
	fmt.Println("main..")
	go say()
	fmt.Println("end...")
}
