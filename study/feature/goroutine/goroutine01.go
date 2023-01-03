package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello", i)
	wg.Done()
}
func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go hello(i)
	}
	fmt.Println("main")
	//time.Sleep(time.Second)
	wg.Wait()
}
