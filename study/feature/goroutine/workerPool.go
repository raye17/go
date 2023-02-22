package main

import (
	"fmt"
	"math/rand"
)

type job struct {
	id      int
	randNum int
}
type result struct {
	job *job
	sum int
}

func main() {
	jobChan := make(chan *job, 128)
	resultChan := make(chan *result, 128)
	createPool(64, jobChan, resultChan)
	go func(resultChan chan *result) {
		for result := range resultChan {
			fmt.Printf("job id:%v randNum :%v result:%d\n",
				result.job.id, result.job.randNum, result.sum)
		}
	}(resultChan)
	var id int
	for {
		id++
		r_num := rand.Int()
		job := &job{
			id:      id,
			randNum: r_num,
		}
		jobChan <- job
	}
}

// 创建工作池
// 参数1，开几个协程
func createPool(num int, jobChan chan *job, resultChan chan *result) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *job, resultChan chan *result) {
			//执行运算
			for job := range jobChan {
				r_num := job.randNum
				var sum int
				for r_num != 0 {
					temp := r_num % 10
					sum += temp
					r_num /= 10
				}
				r := &result{
					job: job,
					sum: sum,
				}
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}
