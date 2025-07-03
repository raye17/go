package main

import "fmt"

type user struct {
	Name string
	Age  int
}

var users []*user

func main() {
	users = []*user{
		{Name: "raye", Age: 18},
		{Name: "sxy", Age: 19},
	}
	for _, v := range users {
		fmt.Println(v)
		fmt.Println(*v)
		fmt.Println(v.Name)
	}
	fmt.Println(">>>>>>>>>>>>>")
	users1 := usersGet()
	for _, v := range users1 {
		fmt.Println(v)
	}
	fmt.Println("<<<<<<<<<<<<<<")

}
func usersGet() (users []user) {
	return users
}
