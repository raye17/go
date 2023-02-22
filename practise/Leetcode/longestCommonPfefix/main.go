package main

import "fmt"

func main() {
	strs1 := []string{"rayeblog", "rayfivdi", "rayecdniv"}
	strs2 := []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(strs1))
	fmt.Println(longestCommonPrefix(strs2))
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return "nil"
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
