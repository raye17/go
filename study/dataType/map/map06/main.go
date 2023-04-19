package main

import (
	"fmt"
)

type Set struct {
	m1 map[string]string
	//m2 map[string]interface{}
	//m3 map[string]int
	m4 map[int]int
}

var set Set

func init() {
	set.m1 = make(map[string]string)
	set.m4 = make(map[int]int)
}
func main() {
	set.m1["str01"] = "str-001"
	set.m4[3] = 4
	for k, v := range set.m1 {
		fmt.Println(k, v)
	}
}
