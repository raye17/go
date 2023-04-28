package main

import "fmt"

type S interface {
	get() string
}
type student struct {
	name string
}

func (s student) get() string {
	return s.name
}

type students []student

func (s students) get() string {
	for _, v := range s {
		v.get()
	}
	return ""
}
func NewStu() student {
	return student{}
}
func main() {
	var gLevel = new(int32)
	fmt.Println(*gLevel)

}
func New(stus ...student) S {
	switch len(stus) {
	case 0:
		return NewStu()
	case 1:
		return stus[0]
	default:
		return students(stus)
	}
}
