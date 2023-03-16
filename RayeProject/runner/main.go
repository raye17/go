package main

import (
	"fmt"
	"log"
	"os"
	"runner/common"
	"time"
)

func main() {
	fmt.Println("start...")
	timeout := time.Second * 5
	r := common.NewRunner(timeout)
	r.Add(common.CreateTask(), common.CreateTask(), common.CreateTask())
	if err := r.Start(); err != nil {
		switch err {
		case common.ErrTimeout:
			log.Println(err)
			os.Exit(1)
		case common.ErrInterrupt:
			log.Println(err)
			os.Exit(2)
		}
	}
	log.Println("over")
}
