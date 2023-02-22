package main

import "fmt"

func main() {
	f1 := func(a, b int) int {
		return a - b
	}
	fmt.Println(f1(3, 7))
	fmt.Println(
		func(x, y int) int {
			return x + y
		}(3, 6))
}
