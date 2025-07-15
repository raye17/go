package main

import (
	"fmt"
	"time"
)

func main() {
	formatTime()
}
func formatTime() {
	var now = time.Now()
	var timestamp = now.Unix()
	year := now.Year()
	month := now.Month()
	fmt.Printf("timeNow:%v\n", now)
	fmt.Printf("year:%v,month:%s\n", year, month)
	fmt.Println("timestamp:", timestamp)
	//
	fmt.Println("*********************")
	fmt.Println(now.Format("2006-01-02 15:04"))
	fmt.Println(now.Format("2006-01-02-15:04:05"))
	fmt.Println(now.Format("2006/01/02 15-04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(now.Format("15:04"))
}
