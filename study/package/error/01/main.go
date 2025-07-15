package main

import (
	"errors"
	"fmt"
)

type user struct {
	name string
	age  int
}

func main() {
	fmt.Println("start")
	var user *user
	user = getUsers()
	go doSomeThing(user)
	fmt.Println("over")
	select {}
}
func createErr() error {
	return errors.New("this is a error")
}
func doSomeThing(user *user) error {
	fmt.Println(user)
	fmt.Println(user.name)
	//return createErr()
	return nil
}
func getUsers() *user {
	var user *user
	return user
}
