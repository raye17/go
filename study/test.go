package main

import "fmt"

func containsDuplicate(nums []int) bool {
	container := make(map[int]int)
	for _, value := range nums {
		if _, ok := container[value]; ok {
			return true
		} else {
			container[value] = 1
		}
	}
	return false
}
func hello(num ...int) {
	num[0] = 199
}
func main() {
	i := []int{1, 2, 3}
	hello(i...)
	fmt.Println(i)
	a := 'a'
	fmt.Println(int(a))
}
func rotate(nums []int, k int) {
	length := len(nums)
	temp := 0
	for i := 0; i < k; i++ {
		temp = nums[length-1]
		for j := length - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = temp
	}
}
