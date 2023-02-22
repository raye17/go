package main

import "fmt"

func main() {
	s := "raye"
	s1 := []byte(s)
	s1[0] = 'R'
	fmt.Println("byte(s1):", s1, "string(s1):", string(s1))
}
