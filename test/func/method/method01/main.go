package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}
func (p *Person) Dream() {
	fmt.Printf("%s的梦想是学好go\n", p.name)
}
func (p *Person) updateAge(age int) {
	p.age = age
}

func (p *Person) show() {
	fmt.Println(*p)
}
func main() {
	p1 := NewPerson("raye", 19)
	p1.Dream()
	fmt.Println(p1.age)
	p1.updateAge(22)
	fmt.Println(p1.age)
	p1.show()
}
