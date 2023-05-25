package main

import "fmt"

type student struct {
	name string
	age  int
}

var m = map[string]*student{
	"001": {
		name: "sxy",
		age:  19,
	},
	"002": {
		name: "sss",
		age:  18,
	},
}

type op struct {
	stu map[string]*student
}

func main() {
	s1 := op{
		stu: m,
	}
	s2 := op{
		stu: m,
	}
	s2.stu["003"] = &student{
		name: "ll",
		age:  100,
	}
	fmt.Println(s1.stu)
}
