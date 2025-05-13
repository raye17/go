package main

import "fmt"

func main() {
	fmt.Println(test4())
}
func test1() (v int) {
	defer fmt.Println(v) // 0
	return v
}
func test2() (v int) {
	defer func() {
		fmt.Println(v) //2
	}()
	return 2
}
func test3() (v int) {
	defer fmt.Println(v) //0
	v = 3
	return 4
}
func test4() (v int) {
	defer func(n int) {
		fmt.Println(n) //0
	}(v)
	return 5
}
func test5() bool {
	a := false
	defer func() { //false
		a = true
	}()
	return a
}
func test6() (a bool) {
	a = false
	defer func() {
		a = true //true
	}()
	return a
}
func test7() (i int) { //i=2
	defer func() {
		i++
		fmt.Println("defer2:", i) // 2.打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 1.打印结果为 defer: 1
	}()
	return i
}
