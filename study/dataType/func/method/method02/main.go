package main

import "fmt"

type User struct {
	id   int
	name string
}

func (u *User) Test() {
	fmt.Printf("%p,%v\n", u, u)
}
func main() {
	u := User{1, "raye"}
	u.Test()
	mValue := u.Test
	mValue() //隐式传递receiver
	mExpression := (*User).Test
	mExpression(&u) //显式传递
}
