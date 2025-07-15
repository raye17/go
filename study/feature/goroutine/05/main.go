package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("start")
	go send()
	wg.Done()
	fmt.Println("over")
	wg.Wait()

}
func send() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		log.Println(i)
	}
}
