package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "raye"
	s2 := "100"
	i1, err := strconv.Atoi(s1)
	check(err)
	fmt.Println("i1: ", i1)
	i2, err := strconv.Atoi(s2)
	check(err)
	fmt.Println("i2: ", i2)
	s3 := strconv.Itoa(i1)
	fmt.Printf("s3: %v,%T\n", s3, s3)
	s4 := strconv.Itoa(i2)
	fmt.Printf("s4: %v,%T\n", s4, s4)
	b, err := strconv.ParseBool("true")
	check(err)
	f, err := strconv.ParseFloat("3.1344", 64)
	check(err)
	i, err := strconv.ParseInt("-76", 10, 64)
	check(err)
	u, err := strconv.ParseUint("8", 0, 64)
	check(err)
	fmt.Println(b, f, i, u)
}
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
