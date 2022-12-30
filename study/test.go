package main

import (
	"flag"
	"fmt"
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
	msg := "raye hello lcx helper"
	s := "hello,孙肖扬"
	for k, v := range s {
		fmt.Printf("k: %v,v: %c\n", k, v)
	}
	n := len(msg)
	fmt.Println(n)
	var length = int32(n)
	fmt.Println(length)
	fmt.Println("***********")
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
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
