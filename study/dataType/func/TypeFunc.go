package main

import "fmt"

type myFunc func(int) int

func (f myFunc) sum(a, b int) int {
	res := a + b
	return f(res)
}
func sum10(num int) int {
	return num * 10

}
func sum100(num int) int {
	return num * 100
}
func handlerSum(handler myFunc, a, b int) int {
	res := handler.sum(a, b)
	fmt.Println(res)
	return res
}
func main() {
	newFunc1 := myFunc(sum10)
	newFunc2 := myFunc(sum100)
	handlerSum(newFunc1, 1, 1)
	handlerSum(newFunc2, 1, 1)
}
