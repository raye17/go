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
	a := 3
	b := 6
	for i := 1; i < 10; i++ {
		if a == i {
			continue
		}

		if b == i {
			break
		}
		fmt.Println(i)
	}
}
