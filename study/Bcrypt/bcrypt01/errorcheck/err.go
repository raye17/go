package errorcheck

import "fmt"

func Check(err error) {
	if err != nil {
		fmt.Println("failed err:", err)
		return
	}
}
