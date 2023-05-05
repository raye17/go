package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["name"] = 11
	m["love"] = 19
	if v, ok := m["name"]; ok {
		fmt.Println(v)
	}
}
