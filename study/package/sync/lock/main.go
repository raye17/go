package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	rwLockT()
}
func lockT() {
	var lock sync.Mutex
	var wg sync.WaitGroup
	var s = "sss"
	lock.Lock()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("go func()")
		lock.Lock()
		fmt.Println("go lock")
		defer lock.Unlock()
		fmt.Println(s)
	}()
	time.Sleep(1 * time.Second)
	s = "sss111"
	lock.Unlock()
	fmt.Println("main unlock")
	wg.Wait()
}
func rwLockT() {
	var lock sync.RWMutex
	var wg sync.WaitGroup
	var s = "sss"
	lock.Lock()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("go func()")
		lock.RLock()
		fmt.Println("go lock")
		defer lock.RUnlock()
		fmt.Println(s)
	}()
	time.Sleep(1 * time.Second)
	s = "sss111"
	lock.Unlock()
	fmt.Println("main unlock")
	wg.Wait()
}
