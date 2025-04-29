package main

import "fmt"

func f() *int {
	v := 1
	fmt.Println(&v)
	return &v
}
func main() {
	var p = f()
	var q = f()
	fmt.Println(p == q)
}
