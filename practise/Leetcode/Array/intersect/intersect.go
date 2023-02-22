package main

import "fmt"

func main() {
	nums1 := []int{4, 9, 5, 4, 4}
	nums2 := []int{9, 4, 9, 8, 4, 6}
	fmt.Println(intersect(nums1, nums2))
}
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := map[int]int{}
	for _, v := range nums1 {
		m[v]++
	}
	var ret []int
	for _, v := range nums2 {
		if m[v] > 0 {

			ret = append(ret, v)
			m[v]--
		}
	}
	return ret
}
