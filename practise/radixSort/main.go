package main

import (
	"fmt"
	"math"
)

func main() {
	s := []int{23, 45, 29, 120, 122, 21, 48}
	RadixSort(s, 0, len(s)-1, 3)
	fmt.Println(s)
}
func RadixSort(arr []int, l, r, digit int) {
	const radix = 10
	bucket := make([]int, r-l+1)
	for d := 1; d <= digit; d++ {
		count := make([]int, radix)
		for i := l; i <= r; i++ {
			j := getDigit(arr[i], d)
			count[j]++
		}
		for i := 1; i < radix; i++ {
			count[i] = count[i] + count[i-1]
		}
		for i := r; i >= l; i-- {
			j := getDigit(arr[i], d)
			bucket[count[j]-1] = arr[i]
			count[j]--
		}
		for i, j := l, 0; i <= r; i, j = i+1, j+1 {
			arr[i] = bucket[j]
		}
	}
}
func getDigit(x, d int) int {
	return x / ((int)(math.Pow(10, float64(d-1)))) % 10
}
