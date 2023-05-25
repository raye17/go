package main

import (
	"fmt"
	"time"
)

// 利用异或进行交换
func main() {
	a, b := 1, 0
	c, d := 11, 9
	t1 := time.Now()
	a = a ^ b
	b = a ^ b
	a = a ^ b
	con1 := time.Since(t1).Microseconds()
	fmt.Println("a:", a, "b:", b, con1)
	t2 := time.Now()
	c, d = d, c
	con2 := time.Since(t2).Microseconds()
	fmt.Println("c:", c, "d:", d, con2)
}
func xor(a, b int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
}
func swap(a, b int) {
	a, b = b, a
}
