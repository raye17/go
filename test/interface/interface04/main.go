package main

import "fmt"

// 定义一个接口
type people interface {
	returnName() string
}

// 定义一个结构体
type student struct {
	name string
}

// 定义student结构体的一个方法
func (s *student) returnName() string {
	return s.name
}
func checkPeople(test interface{}) {
	if _, ok := test.(people); ok {
		fmt.Println("student success people!")
	}
}
func main() {
	stu := &student{name: "raye"}
	var p people = stu
	var p1 people = &student{name: "lcx"}
	fmt.Println(stu.returnName(), p.returnName())
	fmt.Println(p1.returnName())
	fmt.Println("p:", p)
	fmt.Printf("type of p :%T\n", p)
	fmt.Println("stu:", stu)
	fmt.Println("p1:", p1)
	checkPeople(stu)

}
