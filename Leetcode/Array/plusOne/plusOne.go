package main

import "fmt"

func main() {
	e := []int{1, 2, 3}
	fmt.Println(plusOne(e))

}
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	return append([]int{1}, digits...)
}
