package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

var queueName = "q1"
var amqpConn *amqp.Connection

func initLogger() {
	cfg := zap.NewDevelopmentConfig() // 或 zap.NewProductionConfig()
	cfg.DisableStacktrace = true
	logger, _ := cfg.Build()
	zap.ReplaceGlobals(logger) // 替换全局 zap.L()
}

func main() {
	conn, err := amqp.Dial("amqp://root:123456@localhost:5672/sxy")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	amqpConn = conn
	initLogger()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	timeout := time.After(10 * time.Second)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(12 * time.Second)
		receive()
	}()
	go func() {
		time.Sleep(15 * time.Second)
		wg.Done()
	}()
loop:
	for {
		select {
		case <-ticker.C:
			send()
		case <-timeout:
			break loop
		}
	}
	zap.L().Error("test err:", zap.Error(errors.New("test")))
	wg.Wait()
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func send() {
	ch, err := amqpConn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body := "测试111" + " " + time.Now().Format("2006-01-02 15:04:05")
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [producer] Sent %s", body)
}
func receive() {
	ch, err := amqpConn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// q, err := ch.QueueDeclare(
	// 	queueName, // name
	// 	false,     // durable
	// 	false,     // delete when unused
	// 	false,     // exclusive
	// 	false,     // no-wait
	// 	nil,       // arguments
	// )
	// failOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")
	connCloseChan := make(chan *amqp.Error)
	amqpConn.NotifyClose(connCloseChan)

	chCloseChan := make(chan *amqp.Error)
	ch.NotifyClose(chCloseChan)

	go func() {
		for {
			select {
			case err, ok := <-connCloseChan:
				if !ok {
					log.Println("connection close notify channel closed")
					return
				}
				if err != nil {
					log.Println("connection closed:", err)
				}
			case err, ok := <-chCloseChan:
				if !ok {
					fmt.Println("channel close notify channel closed")
					return
				}
				if err != nil {
					fmt.Println("channel closed:", err)
				}
			}
		}
	}()

	for msg := range msgs {
		log.Printf("Received a message: %s", msg.Body)
	}
}
