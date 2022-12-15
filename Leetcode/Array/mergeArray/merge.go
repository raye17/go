package main

import "fmt"

func main() {
	var n1, n2 []int
	n1 = []int{2, 3, 5, 8, 10, 0, 0, 0, 0}
	n2 = []int{1, 4, 8, 15}
	merge(n1, 5, n2, 4)
	fmt.Println(n1)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	for n1, n2, k := m-1, n-1, m+n-1; n1 >= 0 || n2 >= 0; k-- {
		var temp int
		if n1 == -1 {
			temp = nums2[n2]
			n2--
		} else if n2 == -1 {
			temp = nums1[n1]
			n1--
		} else if nums1[n1] > nums2[n2] {
			temp = nums1[n1]
			n1--
		} else {
			temp = nums2[n2]
			n2--
		}
		nums1[k] = temp
	}
}
