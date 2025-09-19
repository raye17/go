package main

import "fmt"

func main() {
	m := make(map[uint64]string)
	//m[0] = "team"
	m[1] = "001"
	m[2] = "002"
	fmt.Println(m)
}
