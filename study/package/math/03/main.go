package main

import (
	"fmt"
	"strconv"
)

func main() {
	amount := float64(100)
	sunAmount := float64(1000)
	rate := strconv.FormatFloat((amount / sunAmount), 'f', -1, 64)
	fmt.Println(rate)
}
