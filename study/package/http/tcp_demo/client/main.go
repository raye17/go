package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//与服务端进行连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("dial failed err:%v\n", err)
		return
	} else {
		fmt.Println("success dial server!")
	}
	//利用该连接进行数据接收和发送
	for i := 0; i < 10; i++ {
		msg := `hello,raye.how are you?`
		conn.Write([]byte(msg))
		//从服务端接收回复的消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed,err:%v\n", err)
			return
		}
		fmt.Println("收到服务端回复: ", string(buf[:n]))
	}
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}
		//给服务端发消息
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("send failed,err:%v\n", err)
			return
		}
		//从服务端接收回复的消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed,err:%v\n", err)
			return
		}
		fmt.Println("收到服务端回复: ", string(buf[:n]))
	}

}
