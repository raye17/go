package main

import (
	"fmt"
	calc "study/packageTest/pack"
)

func main() {
	var x = 4
	var y = 5
	var z = calc.Add(x, y)
	fmt.Println(z)
}
