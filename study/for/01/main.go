package main

import "fmt"

type User struct {
	Name string
	Age  int
	Male bool
}

func main() {
	var users = []*User{
		{
			Name: "张三",
			Male: true,
		},
		{
			Name: "李四",
			Male: true,
		},
		{
			Name: "王五",
			Male: true,
		},
	}
	for _, v := range users {
		fmt.Println(*v)
	}
	i := 0
	for _, v := range users {
		v.Age = 18 + i
		i++
	}
	fmt.Println("*************")
	for _, v := range users {
		fmt.Println(*v)
	}
}
