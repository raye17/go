package main

import "sync"

// 并发资源竞争
// go build -race 可用于查看并发资源竞争
import (
	"fmt"
	"runtime"
)

var (
	count int
	wg    sync.WaitGroup
	mu    sync.Mutex
)

func main() {
	count = 4
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}
func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		mu.Lock()
		value := count
		runtime.Gosched() //让出当前时间片，等待重新分配
		value++
		count = value
		mu.Unlock()
	}
}
