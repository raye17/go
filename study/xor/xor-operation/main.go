package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 1, 7, 2, 4, 5, 6, 7, 9}
	r := []int{0, 0, 0}
	for i := 0; i < len(a); i++ {
		r[0] = r[0] ^ a[i]
	}
	r[1] = r[0] & (^r[0] + 1)
	for i := 0; i < len(a); i++ {
		if a[i]&r[1] == 0 {
			r[2] ^= a[i]
		}
	}
	r[1] = r[0] ^ r[2]
	fmt.Println(r)

}
