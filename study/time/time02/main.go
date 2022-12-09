package main

import (
	"fmt"
	"time"
)

func main() {
	l, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.Now().In(l))
}
