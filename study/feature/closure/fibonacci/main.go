package main

import (
	"fmt"
)

func fibonacci() func() int {
	x, y := 0, 1
	return func() (z int) {
		z, x, y = x, y, x+y
		return
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}
}
