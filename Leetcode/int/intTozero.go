package main

import "fmt"

func main() {
	ret := numberOfSteps(14)
	fmt.Println(ret)
}
func numberOfSteps(num int) int {
	var count = 0
	for num != 0 {
		if num%2 == 0 {
			num /= 2
			count++
		} else {
			num -= 1
			count++
		}
	}
	return count
}
