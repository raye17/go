package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwLock sync.RWMutex
)

func write(i int) {
	rwLock.Lock()
	fmt.Println("已加写锁，在写")
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	fmt.Println("i", i)
	rwLock.Unlock()
	fmt.Println("已解写锁")
	wg.Done()
}
func read(i int) {
	rwLock.RLock()
	fmt.Println("加读锁")
	time.Sleep(time.Millisecond)
	fmt.Println("j: ", i, "x:", x)
	rwLock.RUnlock()
	fmt.Println("解读锁")
	wg.Done()
}
func main() {
	start := time.Now()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go write(i)
	}
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go read(j)
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
