package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Type int    `json:"type"`
	Data []byte `json:"data"`
}

func wsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	log.Println("connect success")
	ticker := time.NewTicker(8 * time.Second)
	defer ticker.Stop()
	msg := make(chan Message)
	go func() {
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				return
			}
			msg <- Message{
				Type: messageType,
				Data: []byte(time.Now().Format("2006-01-02 15:04:05") + string(p)),
			}
		}
	}()
	for {
		select {
		case <-ticker.C:
			if err := conn.WriteMessage(websocket.TextMessage, []byte(time.Now().Format("2006-01-02 15:04:05")+"hello")); err != nil {
				return
			}
		case m := <-msg:
			log.Println("type: ", m.Type)
			if err := conn.WriteMessage(m.Type, m.Data); err != nil {
				return
			}
		}
	}
}
func main() {
	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))
	r.GET("/ws", wsHandler)
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	//http.Handle("/", http.FileServer(http.Dir("./static"))) // 服务静态前端页面

	fmt.Println("WebSocket 服务器启动于 :8080")
	log.Fatal(r.Run(":8080"))
}
