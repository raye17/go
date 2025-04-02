package main

import "fmt"

func Slice01() {
	s1 := make([]int, 5)
	s2 := make([]int, 0)
	fmt.Println(s1, s2)
	s1 = append(s1, 1, 2, 3)
	s2 = append(s2, 12, 3)
	s1 = append(s1, s2...)
	fmt.Println(s1, s2)
}
