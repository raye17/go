package ssh

import (
	"fmt"
	"io"
	"net"
	"ssh/demo/config"
	"time"
)

// GetTunnelConnection 建立并返回SSH隧道连接
func GetTunnelConnection() (net.Listener, error) {
	// 建立SSH客户端连接
	sshClient, err := GetSSHClient()
	if err != nil {
		return nil, fmt.Errorf("SSH连接失败: %w", err)
	}

	// 将远程MySQL端口映射到本地
	localAddr := "127.0.0.1:0" // 0 表示动态分配本地端口
	remoteAddr := fmt.Sprintf("%s:%d", config.AppConfig.MySQL.Host, config.AppConfig.MySQL.Port)
	listener, err := net.Listen("tcp", localAddr)
	if err != nil {
		return nil, fmt.Errorf("本地监听失败: %w", err)
	}

	// 在goroutine中处理隧道转发
	stopChan := make(chan struct{})
	go func() {
		for {
			select {
			case <-stopChan:
				fmt.Println("停止监听")
				return
			default:
				localConn, err := listener.Accept()
				if err != nil {
					fmt.Printf("接受本地连接失败: %v\n", err)
					return
				}
				remoteConn, err := sshClient.Dial("tcp", remoteAddr)
				if err != nil {
					fmt.Printf("SSH隧道连接MySQL失败: %v\n", err)
					localConn.Close()
					return
				}
				go handleTunnel(localConn, remoteConn)
			}
		}
	}()

	// 等待隧道准备就绪
	time.Sleep(1 * time.Second)

	return listener, nil
}

// handleTunnel 处理隧道数据转发
func handleTunnel(localConn, remoteConn net.Conn) {
	defer localConn.Close()
	defer remoteConn.Close()

	// 双向数据转发
	go func() {
		_, err := io.Copy(remoteConn, localConn)
		if err != nil {
			fmt.Printf("本地到远程转发失败: %v\n", err)
		}
	}()
	_, err := io.Copy(localConn, remoteConn)
	if err != nil {
		fmt.Printf("远程到本地转发失败: %v\n", err)
	}
}
