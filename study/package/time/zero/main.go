package main

import (
	"fmt"
	"time"
)

func main() {
	var t time.Time
	fmt.Println(t)
	fmt.Println(t.IsZero())
	t, err := time.Parse(time.RFC3339, "0001-01-01T08:00:00+08:00")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t.IsZero())

}
