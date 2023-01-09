package main

import (
	"fmt"
	"io"
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
func hello(num ...int) {
	num[0] = 199
}

type direction int

const (
	North direction = iota
	east
	south
	west
)

func (d direction) String() string {
	return [...]string{"North", "east", "south", "west"}[d]
}

type name io.Reader

func main() {
	fmt.Println(south)
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
