package main

import (
	"asynq/demo/task"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	t1, err := task.NewEmailSendTask(2222, "EMAIL11", "raye")
	if err != nil {
		panic(err)
	}
	t2, err := task.NewEmailRecvTask(1111, "EMAIL", "lalalla")
	if err != nil {
		panic(err)
	}
	info, err := client.Enqueue(t1, asynq.ProcessIn(12*time.Second), asynq.Retention(3*time.Minute), asynq.Queue("high"))
	if err != nil {
		log.Println("err: ", err)
		panic(err)
	}
	log.Printf("successfuly enqueued task: %+v", info.ID)
	info, err = client.Enqueue(t2, asynq.ProcessIn(1*time.Minute))
	if err != nil {
		panic(err)
	}
	log.Printf("success enqueued task: %+v", info.ID)
}
