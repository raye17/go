package main

import "fmt"

func a() {
	fmt.Println("a")
}
func b() {
	fmt.Println("before panic ")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func b err")
		}
	}()
	panic("panic on b")
}
func c() {
	fmt.Println("c")
}
func main() {
	a()
	b()
	c()

}
