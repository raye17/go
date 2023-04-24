package main

import (
	"fmt"
	"math/rand"
)

const NUM = 17

func main() {
	var num = rand.Intn(100) + 1
	for i := 1; ; i++ {
		if NUM == num {
			fmt.Printf("over,count is %d,num is %d", i, num)
			break
		} else {
			if NUM > num {
				fmt.Printf("%d is littler than NUM\n", num)
				num = rand.Intn(100) + 1
			} else {
				fmt.Printf("%d is bigger than NUM\n", num)
				num = rand.Intn(100) + 1
			}
		}

	}
}
