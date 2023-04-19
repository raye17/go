package main

import (
	"fmt"
)

func main() {
	var (
		i interface{}
		j interface{}
		k interface{}
	)
	i = 3
	j = "str"
	k = 'k'
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["1"] = i
	m["2"] = j
	m["3"] = k
	for k, v := range m {
		switch v.(type) {
		case string:
			fmt.Println("string:", k, v)
		case int:
			fmt.Println("int", k, v)
		case byte:
			fmt.Println("byte", k, v)
		case rune:
			fmt.Println("rune", k, v)
		case nil:
			fmt.Println("nil")
		}
	}
}
