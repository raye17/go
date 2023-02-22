package main

import (
	"bcrypt/hash"
	"fmt"
)

func main() {
	var password = "12345"
	fmt.Println(hash.PasswordHash(password))
}
