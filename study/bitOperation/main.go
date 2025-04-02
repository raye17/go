package main

import "fmt"

func main() {
	n := 8
	suf := make([]int, n+1)
	for i := 0; i < 10; i++ {
		fmt.Println(suf[n] | i)
		fmt.Println("<<")
		fmt.Println(i<<2, i<<3, i<<4)
		fmt.Println("..............")
	}
}
