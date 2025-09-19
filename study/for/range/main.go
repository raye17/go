package main

import (
	"fmt"
	"strconv"
)

func main() {
	var m = make(map[string]string)
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("over")
	fmt.Println(strconv.FormatUint(108, 10))
}
