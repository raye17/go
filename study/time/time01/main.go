package main

import (
	"fmt"
	"time"
)

func main() {
	var now = time.Now()
	var timestamp = now.Unix()
	var times = time.Unix(timestamp, 22)
	var latertime = now.Add(time.Hour)
	fmt.Println(times.Equal(now))
	fmt.Println(now)
	fmt.Println(timestamp)
	fmt.Println(times)
	fmt.Println(latertime)
	fmt.Println(time.Since(now))
}
