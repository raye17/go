package main

import "fmt"

//嵌套匿名结构体字段冲突

type Address struct {
	Province   string
	City       string
	UpdateTime string
}
type Email struct {
	Addr       string
	UpdateTime string
}
type Person struct {
	Name   string
	Gender string
	Age    int
	Address
	Email
}

func main() {
	p1 := Person{
		Name:   "raye",
		Gender: "男",
		Age:    18,
		Address: Address{
			Province:   "河南",
			City:       "周口",
			UpdateTime: "18",
		},
		Email: Email{
			Addr:       "河南",
			UpdateTime: "20",
		},
	}
	fmt.Println(p1)
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.Address)
	//fmt.Println(p1.UpdateTime)
	fmt.Println(p1.Address.UpdateTime)
}
