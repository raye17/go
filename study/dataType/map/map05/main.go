package main

import "fmt"

func main() {
	m1 := make(map[string]string, 1)
	m2 := make(map[int]int, 2)
	m3 := map[string]int{"m3-01": 1, "m3-02": 2}
	fmt.Println(m1, len(m1))
	fmt.Println(m2, len(m2))
	m1["001"] = "m1-01"
	m1["002"] = "m1-02"
	m1["000"] = "m1-00"
	m2[19] = 91
	m2[18] = 81
	m1["003"] = "m1-03"
	fmt.Println(m1, len(m1))
	prints(m1)
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	for k, v := range m3 {
		fmt.Println(k, v)
	}
}
func prints(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
