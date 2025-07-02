package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("这是一个错误")
	fmt.Println(err)
	if err := Add(-1, 2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
}
func Add(a, b int) error {
	if a < 0 || b < 0 {
		return errors.New("a和b不能小于0")
	}
	return nil
}
