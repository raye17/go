package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	err := c.AddFunc("*/30 * * * * *", func() {
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println("time: ", now, "hello")
	})
	//fmt.Println(id)
	if err != nil {
		panic(err)
	}
	fmt.Println("test")
	c.Start()
	select {}
}
