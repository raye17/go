package main

import (
	"fmt"
	"raye17/sort/quickSort"
)

func main() {
	s := []int{5, 2, 7, 9, 4, 16, 10, 3, 21, 15, 32, 18, 1}
	quickSort.QuickSort(s, 0, len(s)-1)
	fmt.Println("s:", s)
}
