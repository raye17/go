package main

import "fmt"

func main() {
	var i interface{} = "100"
	fmt.Printf("%v,%#v,%T\n", i, i, i)
	i = 100
	fmt.Printf("%v,%#v,%T\n", i, i, i)
	i = true
	fmt.Printf("%v,%#v,%T", i, i, i)

}
