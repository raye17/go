package test

import "fmt"

var i = 4

const name = "test"

func Test() {
	fmt.Println("Test...")
}

func init() {
	fmt.Println("init...")
	fmt.Println(i)
	fmt.Println(name)
}
