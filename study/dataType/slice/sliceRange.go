package main

//for range 循环的时候会创建每个元素的副本，而不是元素的引用
import "fmt"

func Slice02() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for key, val := range slice {
		value := val
		m[key] = &value
	}
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
