package main

import (
	"fmt"
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

type inter01 interface {
	get() string
}
type str01 struct {
	inter01
	name string
}

func (s *str01) get() string {
	ss := s.inter01.get()
	fmt.Println(ss)
	return ss
}
func main() {
	s := str01{
		name: "string",
	}
	s.get()
}
