package main

import "fmt"

//封装及绑定

type Person struct {
	name   string
	age    int
	gender string
}

func (p Person) eat() {
	p.name = "ray"
	fmt.Println("不使用指针：", p.name)
}
func (p *Person) age1() {
	p.age = 1000
	fmt.Println("使用指针:", p.age)
}
func main() {
	ppl := Person{
		name:   "test",
		age:    19,
		gender: "男",
	}
	ppl.eat()
	fmt.Println(ppl.name)
	ppl.age1()
	fmt.Println(ppl.age)
}
