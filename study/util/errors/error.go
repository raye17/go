package errors

import "fmt"

func Checkout(msg string, err error) {
	if err != nil {
		fmt.Println(msg, ":", err)
	}
}
