package main

import "fmt"

func add(a ...int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}
func main() {
	fmt.Println(add(10, 20, 30, 40))
}
