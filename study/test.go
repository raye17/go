package main

import (
	"fmt"
	"reflect"
)

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
	var chan01 = make(chan int)
	var chan02 = make(chan int)
	fmt.Println(reflect.ValueOf(chan01))
	go func() {
		x := <-chan02
		fmt.Println("x:", x)
	}()
	chan02 <- 3

	fmt.Println(chan01)
	//chan02 <- 4
	fmt.Println(reflect.ValueOf(chan02))
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
