package main

import (
	"fmt"
)

func main() {
	//fmt.Println(test1())
	//fmt.Println(test2())
	//fmt.Println(test3())
	//fmt.Println(test4())
	//fmt.Println(test7())
	s := make([]int, 3)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

}
func test7() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i
}
func test5() bool {
	a := false
	defer func() {
		a = true
	}()
	return a
}
func test6() (a bool) {
	a = false
	defer func() {
		a = true
	}()
	return a
}
func test1() (v int) {
	defer fmt.Println(v)
	return v
}
func test2() (v int) {
	defer func() {
		fmt.Println(v)
	}()
	return 2
}
func test3() (v int) {
	defer fmt.Println(v)
	v = 3
	return 4
}
func test4() (v int) {
	defer func(n int) {
		fmt.Println(n)
	}(v)
	return 5
}
