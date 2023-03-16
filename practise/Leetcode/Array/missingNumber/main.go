package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 6, 7, 8, 4}
	r := missingNumber(nums)
	fmt.Println(r)
}

// 给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数
func missingNumber(nums []int) int {
	//sum01, sum02 := 0, 0
	//for i := 0; i <= len(nums); i++ {
	//	sum01 += i
	//}
	//for _, v := range nums {
	//	sum02 += v
	//}
	//return sum01 - sum02
	//异或
	result := 0
	for k, v := range nums {
		result ^= k ^ v
	}
	return result ^ len(nums)
}
