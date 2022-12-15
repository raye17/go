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

func main() {
	for i := 1; i < 15; i++ {
		if i%2 == 0 {
			fmt.Println("偶数")
			continue
		}
		if i == 9 {
			break
		}
		fmt.Println("i:", i)
	}
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
