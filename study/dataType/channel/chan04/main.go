package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//开启三个goroutine
	for i := 1; i < 4; i++ {
		go worker(i, jobs, results)
	}
	//发送五个任务
	for i := 1; i < 6; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 1; i < 6; i++ {
		ret := <-results
		fmt.Println(ret)
	}
}
