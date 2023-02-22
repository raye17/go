package main

import (
	"fmt"
	"time"
)

func main() {
	var now = time.Now()
	var timestamp = now.Unix()
	year := now.Year()
	month := now.Month()
	fmt.Printf("timeNow:%v\n", now)
	fmt.Printf("year:%v,month:%s\n", year, month)
	fmt.Println("timestamp:", timestamp)
	var times = time.Unix(timestamp, 0)
	var latertime = now.Add(time.Hour)
	fmt.Println(times.Equal(now))
	fmt.Println(times)
	fmt.Println("after one hour:", latertime)
	fmt.Println(time.Since(now))
	fmt.Println(latertime.After(now))
	fmt.Println("time.Tick")
	//tickDemo()
	//时区
	fmt.Println("时区")
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	time, err := time.ParseInLocation("2006/01/02 15:04:05 Mon Jan", "2006/01/02 15:04:05 Mon Jan", loc)
	fmt.Println(time)
}
func tickDemo() {
	ticker := time.Tick(time.Second * 5)
	for i := range ticker {
		fmt.Println(i.Format("2006-01-02 15:04:05 Mon Jan"))
	}
}
