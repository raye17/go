package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
		fmt.Println("not break")
	}
}
