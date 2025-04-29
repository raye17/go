package mq

import (
	"encoding/json"
	"fmt"
	"log"
	"raye/demo/db/model"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel

func init() {
	var err error
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	fmt.Println("mq init success")
}

func PushMsg(user model.User) error {
	fmt.Println("push msg...")
	// 发布消息到RabbitMQ
	body, err := json.Marshal(user)
	if err != nil {
		log.Printf("Failed to marshal user data: %v", err)
		return err
	}
	err = ch.Publish(
		"",
		"user_created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish a message: %v", err)
		return err
	}
	fmt.Println("push msg success")
	return nil
}
func Listening() error {
	err := ConsumeUserCreated()
	if err != nil {
		return err
	}
	return nil
}
func ConsumeUserCreated() error {
	// 声明队列
	q, err := ch.QueueDeclare(
		"user_created", // 队列名
		false,          // 是否持久化
		false,          // 是否自动删除
		false,          // 是否排他
		false,          // 额外参数
		amqp.Table{},
	)
	if err != nil {
		log.Printf("Failed to declare a queue: %v", err)
		return err
	}

	// 消费消息
	msgs, err := ch.Consume(
		q.Name, // 队列名
		"",     // 消费者标识
		true,   // 自动应答
		false,  // 是否排他
		false,  // 是否阻塞
		false,  // 额外参数
		nil,    // 额外参数

	)
	if err != nil {
		log.Printf("Failed to register a consumer: %v", err)
		return err
	}

	// 处理消息
	go func() {
		for d := range msgs {
			var user model.User
			if err := json.Unmarshal(d.Body, &user); err != nil {
				log.Printf("Failed to unmarshal user data: %v", err)
				continue
			}
			log.Printf("Received a user: %+v", user)
		}
	}()

	log.Println("Waiting for messages. To exit press CTRL+C")
	return nil
}
