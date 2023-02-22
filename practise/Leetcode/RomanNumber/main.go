package main

import "fmt"

/*
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/
func main() {
	s1 := "III"
	s2 := "LVIII"
	s3 := "MCMXCIV"
	for i := 0; i < 3; i++ {
		switch i {
		case 0:
			ret := romanToInt(s1)
			fmt.Println(ret)
		case 1:
			ret := romanToInt(s2)
			fmt.Println(ret)
		case 2:
			ret := romanToInt(s3)
			fmt.Println(ret)
		}

	}

}

func romanToInt(s string) int {
	var romanNumber = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	ret := 0
	for i := range s {
		value := romanNumber[s[i]]
		if i < len(s)-1 && value < romanNumber[s[i+1]] {
			ret -= value
		} else {
			ret += value
		}
	}
	return ret
}
