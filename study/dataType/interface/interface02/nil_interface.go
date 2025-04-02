package main

import "fmt"

func Main1() {
	var i interface{} = "100"
	fmt.Printf("%v,%#v,%T\n", i, i, i)
	i = 100
	fmt.Printf("%v,%#v,%T\n", i, i, i)
	i = true
	fmt.Printf("%v,%#v,%T", i, i, i)

}
