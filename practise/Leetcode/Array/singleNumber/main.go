package main

import "fmt"

func main() {
	nums01 := []int{1, 3, 5, 7, 9, 7, 5, 3, 1}
	r := singleNumber(nums01)
	fmt.Println(r)
}

func singleNumber(nums []int) int {
	//一个非空整数数组nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
	//设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
	sum := 0
	for _, v := range nums {
		sum ^= v
	}
	return sum
}
