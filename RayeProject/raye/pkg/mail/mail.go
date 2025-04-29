package mail

import (
	"fmt"
	"net/smtp"
	"raye/demo/config"
	"strings"

	"github.com/streadway/amqp"
)

// EmailTemplate 邮件模板结构体
type EmailTemplate struct {
	Subject string
	Body    string
}

// SendEmail 发送邮件
func SendEmail(to []string, subject string, body string) error {
	cfg := config.AppConfig.SMTP
	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", strings.Join(to, ","), subject, body))

	err := smtp.SendMail(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), auth, cfg.From, to, msg)
	if err != nil {
		return fmt.Errorf("发送邮件失败: %v", err)
	}
	return nil
}

// SendTemplateEmail 发送模板邮件
func SendTemplateEmail(to []string, template EmailTemplate, data map[string]string) error {
	subject := template.Subject
	body := template.Body

	for k, v := range data {
		subject = strings.ReplaceAll(subject, "{{"+k+"}}", v)
		body = strings.ReplaceAll(body, "{{"+k+"}}", v)
	}

	return SendEmail(to, subject, body)
}

// ListenForEmailTasks 监听RabbitMQ邮件发送任务
func ListenForEmailTasks(queue string) error {
	conn, err := amqp.Dial(config.AppConfig.RabbitMQ.URL)
	if err != nil {
		return fmt.Errorf("连接RabbitMQ失败: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("创建通道失败: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("声明队列失败: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return fmt.Errorf("注册消费者失败: %v", err)
	}

	for d := range msgs {
		// 这里处理邮件发送任务
		fmt.Printf("收到邮件发送任务: %s\n", d.Body)
		// 需要根据实际业务需求解析消息内容
	}

	return nil
}
