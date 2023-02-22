package main

import "fmt"

// 函数作为参数传递
type calcFunc func(x, y int) int

func addFunc(a, b int) int {
	return a + b
}
func subFunc(a, b int) int {
	return a - b
}
func operationFunc(x, y int, calcFunc calcFunc) int {
	return calcFunc(x, y)
}
func main() {
	sum := operationFunc(20, 19, addFunc)
	sub01 := operationFunc(23, 19, subFunc)
	fmt.Println(sum, sub01)
}
