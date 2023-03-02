package main

import (
	"fmt"
)

func main() {
	fmt.Println("start...")
	i := 1
add:
	if i < 5 {
		i++
		fmt.Println(i)
		goto add
	}
}
