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
	var test = map[int]byte{
		1: '{',
		2: '[',
	}
	if test[3] == 0 {
		fmt.Println("exit")
	} else {
		fmt.Println("not exist")
	}
	value, ok := test[3]
	fmt.Println(value, ok)

}
