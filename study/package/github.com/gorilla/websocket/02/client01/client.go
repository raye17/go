package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	serverUrl     = "ws://localhost:8080/ws?connId=%s"
	maxRetryTimes = 10
	retryDelay    = 1 * time.Second
)

var ErrReconnectNeeded = errors.New("connection lost, need reconnect")

func main() {
	uuid := uuid.New().String()
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf(serverUrl, uuid), nil)
	if err != nil {
		log.Fatal("connect failed")
		conn = connectWithRetry()
		return
	}
	log.Println("connect server success")
	for {
		err := data(conn)
		if err == ErrReconnectNeeded {
			conn = connectWithRetry()
			if conn == nil {
				log.Fatal("connect failed")
				return
			}
			log.Println("connect retry server success")
			continue
		} else if err == nil {
			log.Println("exit")
			break
		} else {
			log.Println("read error:", err)
			break
		}

	}
}
func data(conn *websocket.Conn) (err error) {
	ticker := time.NewTicker(5 * time.Second)
	defer conn.Close()
	defer ticker.Stop()
	errChan := make(chan error, 1)
	// 启动一个协程监听服务器消息
	go func() {
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				errChan <- ErrReconnectNeeded
				return
			}
			if messageType == websocket.TextMessage {
				log.Println("⬅️ 收到文本消息:", string(message))
			} else {
				log.Printf("⬅️ 收到非文本消息，类型：%d\n", messageType)
			}
		}
	}()
	go func() {
		for {
			select {
			case <-ticker.C:
				err = conn.WriteMessage(websocket.PingMessage, []byte("ping"))
				if err != nil {
					log.Println("write ping message error:", err)
					errChan <- ErrReconnectNeeded
					return
				} else {
					//log.Println("write ping message success")
				}
			case <-errChan:
				return
			}
		}
	}()
	time.Sleep(100 * time.Millisecond)
	// 主协程等待用户输入，发送消息给服务器
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("请输入消息，按回车发送，输入 exit 退出：")
		fmt.Print("> ")
		for scanner.Scan() {
			text := scanner.Text()
			if text == "exit" {
				fmt.Println("退出客户端")
				errChan <- ErrReconnectNeeded
				return
			}
			err = conn.WriteMessage(websocket.TextMessage, []byte(text))
			if err != nil {
				return
			}
			fmt.Println("消息已发送")
			time.Sleep(100 * time.Millisecond)
			fmt.Print("> ")
		}
	}()
	return <-errChan
}
func connectWithRetry() *websocket.Conn {
	for i := 1; i <= maxRetryTimes; i++ {
		log.Println("connecting retry...")
		fmt.Printf("🔄 尝试第 %d 次连接...\n", i)
		conn, _, err := websocket.DefaultDialer.Dial(serverUrl, nil)
		if err == nil {
			log.Printf("✅ 第 %d 次连接成功", i)
			return conn
		}
		log.Printf("第%d次重试连接失败,错误:%v", i, err)
		time.Sleep(retryDelay)
	}
	log.Println("⛔ 重连超过最大次数，退出程序")
	return nil
}
