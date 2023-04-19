package main

import "fmt"

// 可变参数
func main() {
	fmt.Println(add(1, 2, 3, 4, 5))
	fmt.Println(add(4, 5, 6))
}
func add(nums ...int) int {
	length := len(nums)
	fmt.Println("length:", length)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
