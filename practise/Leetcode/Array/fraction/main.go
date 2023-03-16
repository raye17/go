package main

import "fmt"

func main() {
	n := []int{3, 2, 0, 2}
	fmt.Println(fraction(n))
}
func fraction(cont []int) []int {
	v1, v2 := f(cont)
	return []int{v2, v1}
}
func f(cont []int) (int, int) {
	if len(cont) == 1 {
		return 1, cont[0]
	}
	v1, v2 := f(cont[1:])
	return v2, cont[0]*v2 + v1
}
