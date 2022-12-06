package main

import (
	"fmt"
	"os"
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
	s, err := os.Getwd()
	if err != nil {
		fmt.Println("")
	}
	fmt.Println(s)
}
