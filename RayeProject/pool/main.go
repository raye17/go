package main

import (
	"sync"
)

const (
	maxGoroutine = 5
	poolRes      = 2
)

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)
}
