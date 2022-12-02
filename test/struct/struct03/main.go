package main

import "fmt"

//结构体的继承

type Animal struct {
	name string
}

func (a Animal) move() {
	fmt.Println(a.name, "会动~")
}

type Dog struct {
	Feet    int
	*Animal //匿名嵌套，结构体指针
}

func (d *Dog) Bark() {
	fmt.Println(d.name, "会汪汪汪~")
}
func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "大黄",
		},
	}
	d1.move()
	d1.Bark()
}
