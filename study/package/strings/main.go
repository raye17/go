package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isContains("sxy", "raye"))
	s := "RayeLi"
	fmt.Println(toLower(s))
}
func isContains(a, b string) bool {
	return strings.Contains(a, b)
}
func toLower(s string) string {
	return strings.ToLower(s)
}
