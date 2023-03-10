package main

import (
	"fmt"
	"os"
)

type i interface {
	get()
	set(string, int)
}
type stu struct {
	name string
	age  int
}

func (s stu) get() {
	fmt.Println(s.name, ":", s.age)
}
func (s stu) set(name string, age int) {
	s.name = name
	s.age = age
}

type teacher struct {
	name string
	age  int
}

func (t *teacher) get() {
	fmt.Println(t.name, t.age)
}
func (t *teacher) set(name string, age int) {
	t.name = name
	t.age = age
}
func main() {
	os.Setenv("name", "raye")
	fmt.Println(os.Getenv("name"))
}
