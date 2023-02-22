package main

import "fmt"

type i interface {
	run()
}
type s struct {
}

func (*s) run() {
	fmt.Println("run...")
}

type person struct {
	s
	name string
	age  int
}

func main() {
	p := person{}
	p.run()
}
