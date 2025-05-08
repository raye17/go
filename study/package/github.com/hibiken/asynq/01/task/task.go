package task

import (
	"asynq/demo/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

const (
	TaskTypeEmailSend = "email:send"
	TaskTypeEmailRecv = "email:recv"
)

func NewEmailSendTask(userId int, email string, name string) (*asynq.Task, error) {
	payload, err := json.Marshal(&utils.EmailTaskPayload{UserId: userId, Email: email, Username: name})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TaskTypeEmailSend, payload), nil
}
func NewEmailRecvTask(userId int, email string, name string) (*asynq.Task, error) {
	payload, err := json.Marshal(&utils.EmailTaskPayload{UserId: userId, Email: email, Username: name})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TaskTypeEmailRecv, payload), nil
}
func HandleSendEmailTask(ctx context.Context, t *asynq.Task) error {
	var p utils.EmailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Println("Sending Email to User:", p)
	return nil
}
func HandleRecvEmailTask(ctx context.Context, t *asynq.Task) error {
	var p utils.EmailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Receiving Email to User: user_id=%d", p.UserId)
	return nil
}
