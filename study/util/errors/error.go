package errors

import "fmt"

func Checkout(s string, err error) {
	if err != nil {
		fmt.Println(s, ":", err)
	}
}
