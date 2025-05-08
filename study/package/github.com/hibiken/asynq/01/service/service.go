package main

import (
	"asynq/demo/task"
	"log"

	"github.com/hibiken/asynq"
)

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: "localhost:6379"},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"high":    6,
				"default": 3,
			},
		},
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TaskTypeEmailSend, task.HandleSendEmailTask)
	mux.HandleFunc(task.TaskTypeEmailRecv, task.HandleRecvEmailTask)
	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
