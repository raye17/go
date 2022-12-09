package main

import "fmt"

func main() {
	s1 := "{}()[]"
	s2 := "{()}"
	ret := isValid(s1)
	fmt.Println(ret)
	ret = isValid(s2)
	fmt.Println(ret)
}
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	paris := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < n; i++ {
		if paris[s[i]] > 0 {
			if len(stack) == 0 || paris[s[i]] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
