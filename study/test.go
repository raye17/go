package main

import (
	"fmt"
	"sync/atomic"
)

var opts int64 = 0

func main() {
	fmt.Printf("%v\n", &opts)
	add(&opts, 3)
}

func add(addr *int64, delta int64) {
	atomic.AddInt64(addr, delta)
	fmt.Println("add opts:", *addr)
}
