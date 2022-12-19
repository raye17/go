package main

import "fmt"

// 结构体初始化

type Person struct {
	name  string
	age   int
	score map[string]*Person
}

func EmptyPerson() *Person {
	return &Person{
		name:  "nil",
		age:   0,
		score: map[string]*Person{},
	}
}
func main() {
	p := Person{}
	p1 := EmptyPerson()
	fmt.Println("name:", p.name, "age:", p.age, p.score)
	fmt.Println(*p1)
	fmt.Println("name:", p1.name, "age:", p1.age, p1.score)
}
