package main

import (
	"fmt"
	calic "test/test03/packageTest/pack"
)

func main() {
	var x = 4
	var y = 5
	var z = calic.Add(x, y)
	fmt.Println(z)
}
