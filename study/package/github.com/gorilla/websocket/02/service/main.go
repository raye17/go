package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()
	r.GET("/ws", wsHandler)
	fmt.Println("WebSocket 服务器启动于 :8080")
	log.Fatal(r.Run(":8080"))
}

type Message struct {
	Type int    `json:"type"`
	Data []byte `json:"data"`
}

var (
	clients     = make(map[string]*websocket.Conn)
	clientsLock sync.Mutex
)

var connStatus bool

func wsHandler(c *gin.Context) {
	connId := c.Query("connId")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	log.Println("connId: ", connId)
	clientsLock.Lock()
	clients[connId] = conn
	clientsLock.Unlock()

	successChan := make(chan Message)
	disconnectChan := make(chan struct{})
	connClosed := make(chan struct{})
	msg := make(chan Message)
	var connStatus = true
	log.Println("client connect success")
	go func() {
		successChan <- Message{
			Type: websocket.TextMessage,
			Data: []byte("connect success"),
		}
	}()
	go func() {
		<-disconnectChan
		log.Printf("connId: %s, 客户端断开,等待10秒后关闭连接...", connId[:4])
		time.Sleep(10 * time.Second)
		if !connStatus {
			close(connClosed)
			log.Printf("connId: %s, 连接关闭", connId[:4])
			conn.Close()
		}
	}()
	ticker := time.NewTicker(800 * time.Second)
	defer ticker.Stop()
	go func() {
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				if connStatus {
					connStatus = false
					close(disconnectChan)
				}
				return
			}
			//log.Println(time.Now().Format("2006-01-02 15:04:05"), "receive msg: ", string(p))
			msg <- Message{
				Type: messageType,
				Data: []byte(string(p)),
			}
		}
	}()
	for {
		select {
		case <-ticker.C:
			if err := conn.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
				connStatus = false
				close(disconnectChan)
				return
			}
		case m := <-msg:
			broadcastMessage(connId, m)
			//log.Printf("connId: %s, type: %d, receive data: %s", connId[:4], m.Type, string(m.Data))
		case successMsg := <-successChan:
			if err := conn.WriteMessage(successMsg.Type, successMsg.Data); err != nil {
				connStatus = false
				close(disconnectChan)
				return
			}
		case <-connClosed:
			return
		}
	}
}
func broadcastMessage(senderId string, msg Message) {
	clientsLock.Lock()
	defer clientsLock.Unlock()

	for id, conn := range clients {
		var sendMsg string
		if id == senderId {
			sendMsg = fmt.Sprintf("你说：%s", msg.Data)
		} else {
			sendMsg = fmt.Sprintf("%s说,%s", senderId[:4], msg.Data)
		}
		err := conn.WriteMessage(msg.Type, []byte(sendMsg))
		if err != nil {
			log.Printf("发送消息给connId %s 失败: %v", id[:4], err)
		}
	}
}
