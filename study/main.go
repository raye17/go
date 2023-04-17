package main

import "fmt"

type student struct {
	name string
	age  int
}

func (s *student) changeName(name string) string {
	s.name = name
	return s.name
}

type func01 func(string) string

func test(s string, func012 func01) {
	fmt.Println(s)
	func012("sss")
}
func main() {
	s := student{
		name: "sxy",
		age:  19,
	}
	test("test", s.changeName)
	fmt.Println(s.name)
}
