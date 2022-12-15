package main

import (
	"fmt"
	"sort"
)

func main() {
	var s = []int{0, 5, 7, 3, 6, 4}
	sort.Ints(s)
	fmt.Println(s)
	for i := 0; i < 6; i = i + 2 {

	}
}
func addToArrayForm(num []int, k int) (ans []int) {
	for i := len(num) - 1; i >= 0; i-- {
		sum := num[i] + k%10
		k /= 10
		if sum >= 10 {
			k++
			sum -= 10
		}
		ans = append(ans, sum)
	}
	for ; k > 0; k /= 10 {
		ans = append(ans, k)
	}
	reverse(ans)
	return
}
func reverse(num []int) {
	for i, n := 0, len(num); i < n/2; i++ {
		num[i], num[n-i-1] = num[n-1-i], num[i]
	}
}
