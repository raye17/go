package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) set(name string, age int) {
	p.name = name
	p.age = age
}

type Student struct {
	*Person
	Sid   int
	class string
}

func main() {
	stu := &Student{
		Person: &Person{
			name: "raye",
			age:  19,
		},
		Sid:   36001,
		class: "36班",
	}
	fmt.Println(*stu.Person)
	stu.set("sxy", 1000)
	fmt.Println(*stu.Person)
}
