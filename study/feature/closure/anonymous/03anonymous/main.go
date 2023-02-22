package main

import "fmt"

// 函数作为返回值传递
func add(x, y int) func() int {
	f := func() int {
		return x + y
	}
	return f
}
func sub(x, y int) func() int {
	return func() int {
		return x - y
	}
}
func mul(x, y int) int {
	return func() int {
		return x * y
	}()
}
func calc01(x, y int) (func(int) int, func() int) {
	f1 := func(z int) int {
		return (x + y) * z / 2
	}
	f2 := func() int {
		return 2 * (x + y)
	}
	return f1, f2
}
func calc02(x, y int) (func() int, func() int, int) {
	f3 := add(x, y)
	f4 := sub(x, y)
	return f3, f4, mul(f3(), f4())
}
func main() {
	r1, r2 := calc01(3, 4)
	fmt.Println(r1(4), r2())
	r3, r4, r5 := calc02(5, 6)
	fmt.Println(r3(), r4(), r5)
	fmt.Println(mul(3, 4))
}
