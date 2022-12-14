package main

import "fmt"

func main() {
	s := "my name is raye"
	ret := reverseWords(s)
	fmt.Println(ret)
}
func reverseWords(s string) string {
	b := []byte(s)
	head := 0
	for tail := 0; tail < len(b); tail++ {
		if b[tail] == ' ' {
			for left, right := head, tail-1; left < right; left, right = left+1, right-1 {
				b[left], b[right] = b[right], b[left]
			}
			head = tail + 1
		}
		if tail == len(b)-1 {
			for left, right := head, tail; left < right; left, right = left+1, right-1 {
				b[left], b[right] = b[right], b[left]
			}
		}
	}
	return string(b)

}
func reverseString(s []byte) {
	n := len(s)
	i, j := 0, n-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
