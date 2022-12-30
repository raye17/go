package main

import (
	"fmt"
)

func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	m["raye"] = s
	s = append(s[:1], s[2:]...)
	fmt.Println(s)
	fmt.Println(m["raye"])
	fmt.Println("-------------------")
	s01 := []int{1, 2, 3, 4, 5}
	s02 := s01[1:4]
	s02[2] = 29
	fmt.Println(s02)
	fmt.Println(s01)
}
