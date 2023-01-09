package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func demo(ce []student) {
	ce[1].Age = 999
}
func main() {
	m := make(map[string]*student)
	stus := []student{
		{1, "raye", 22},
		{2, "sxy", 21},
		{3, "lcx", 23},
	}
	data, _ := json.Marshal(stus)
	fmt.Println(string(data))
	for _, stu := range stus {
		value := stu
		m[stu.Name] = &value
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.Name, v.Age)
	}
	p := make(map[int]student)
	p[1] = student{
		Id:   017,
		Name: "001",
		Age:  100,
	}
	p[2] = student{99, "002", 199}
	fmt.Println(p)
	delete(p, 2)
	fmt.Println(p)
}
