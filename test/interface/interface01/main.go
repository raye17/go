package main

import "fmt"

type cat struct{}
type dog struct{}
type speaker interface {
	speak()
}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}
func da(s speaker) {
	s.speak()
}
func main() {
	var c1 cat
	var d1 dog
	da(c1)
	da(d1)
}
