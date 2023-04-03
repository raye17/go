package main

import "fmt"

func main() {
	//f := incr()
	//for i := 0; i < 3; i++ {
	//	fmt.Println(f())
	//}
	//for i := 0; i < 3; i++ {
	//	fmt.Println(incr()())
	//}
	i := 3
	f := func() {
		fmt.Println(&i)
	}
	f()
	i = 6
	fmt.Println(&i)
}
func incr() func() int {
	fmt.Println("before...")
	var x int
	fmt.Println("before 01..")
	return func() int {
		fmt.Println("func before ...")
		x++
		return x
	}
}
