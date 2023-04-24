package main

import (
	"fmt"
	"math/rand"
)

func main() {
	mm, p := circle(100000000)
	fmt.Println(mm, p)
}
func circle(n int) (m map[string]int, per float64) {
	m = map[string]int{}
	for i := 0; i < n; i++ {
		j := rand.Intn(2)
		switch j {
		case 0:
			m["0"]++
		case 1:
			m["1"]++
		}
	}
	per = float64(m["0"]) / float64(n)
	return m, per
}
