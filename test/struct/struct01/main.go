package main

import "fmt"

//匿名结构体

type Address struct {
	Province string
	City     string
}

type Person struct {
	Name   string
	Gender string
	Age    int
	Address
}

func main() {
	p1 := Person{
		Name:   "raye",
		Gender: "男",
		Age:    18,
		Address: Address{
			Province: "河南",
			City:     "周口",
		},
	}
	fmt.Println(p1)
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.Address)
	fmt.Println(p1.Province)
	fmt.Println(p1.Address.Province)
}
