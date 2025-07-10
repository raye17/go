package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	// for t := range ticker.C {
	// 	fmt.Println(t)
	// }
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case <-c:
			fmt.Println("stop")
			return
		case t := <-ticker.C:
			fmt.Println("ticker: ", t.Format("2006-01-02 15:04:05"))
		}
	}
}
