package main

import . "fmt"

func main() {
	s := []int{1, 2, 3} // 初始底层数组
	Printf("原数组地址: %p\n", &s[0])
	s = append(s, 4, 3, 6, 7, 8, 9, 12) // 容量不足时创建新数组
	Printf("新数组地址: %p\n", &s[0])        // 地址可能改变
}
