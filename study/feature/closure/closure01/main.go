package main

import (
	"fmt"
	"time"
)

// 普通闭包
func add() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

// 无状态，无变量的闭包
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

// 使用闭包实现斐波那契数列
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func main() {
	//普通闭包
	a := add()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+...+%d = %d\n", i, a(i))
	}
	//无状态，无变量的闭包
	b := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, b = b(i)
		fmt.Printf("0+...+%d = %d\n", i, s)
	}
	fmt.Println("fibonacci...")
	f := fibonacci()
	for i := 1; ; i++ {
		time.Sleep(time.Second)
		fmt.Println(f())
	}
	//fmt.Println(f(), f(), f(), f(), f(), f(), f())
}
