package main

import "fmt"

// 结构体的继承

type Animal interface {
	Name() string
	Speak() string
	Play()
}
type Dog struct {
	Gender string
	name   string
}

func (d *Dog) Name() string {
	return d.name
}
func (d *Dog) Play() {
	fmt.Println(d.Speak())
}

func (d *Dog) Speak() string {
	return fmt.Sprintf("%v and my gender is %v", d.name, d.Gender)
}
func Play(a Animal) {
	a.Play()
}
func main() {
	d1 := &Dog{
		name:   "大黄",
		Gender: "Male",
	}
	fmt.Println(d1.Name())
	fmt.Println(d1.Speak())
	Play(d1)
}
