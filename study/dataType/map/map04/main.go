package main

import "fmt"

func main() {
	m1 := make(map[string]string)
	m1["raye"] = "sxy"
	m1["001"] = "sss"
	m2 := make(map[string]string)
	for k, v := range m1 {
		m2[k] = v
	}
	fmt.Printf("%p,%p\n", m1, m2)
	fmt.Println(m2)
}
