package main

import "fmt"

type inter interface {
	update()
}
type emptyInter struct {
	empty string
}

func (e *emptyInter) update() {
	fmt.Println("update", e.empty)
}

type student struct {
	sid   int
	class string
	i     inter
}

func main() {
	stu := &student{
		sid:   36001,
		class: "36班",
		i: &emptyInter{
			empty: "empty",
		},
	}
	stu.i.update()
	fmt.Printf("%+v", *stu)
}
