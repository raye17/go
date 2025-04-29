package main

import (
	"fmt"
	"study/redis/cache"
)

func main() {
	redis := cache.NewClientRedis()
	fmt.Println(redis)
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
