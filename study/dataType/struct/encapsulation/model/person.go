package model

import "fmt"

type person struct {
	Name string
	age  int
	sal  float64
}

// NewPerson 构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) Set(age int, sal float64) {
	if age > 0 && age < 120 {
		p.age = age
	} else {
		fmt.Println("the age is invalid")
	}
	if sal >= 3000 && sal < 10000 {
		p.sal = sal
	} else {
		fmt.Println("sal is invalid")
	}
}
func (p *person) Get() (int, float64) {
	return p.age, p.sal
}
