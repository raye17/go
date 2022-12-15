package main

import (
	"fmt"
	"study/Bcrypt/bcrypt01/hash"
)

func main() {
	var password = "12345"
	fmt.Println(hash.PasswordHash(password))
}
