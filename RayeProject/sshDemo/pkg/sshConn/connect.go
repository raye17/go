package sshconn

import (
	"fmt"
	"io"
	"net"
	"ssh/demo/config"
	"time"

	"golang.org/x/crypto/ssh"
)

const (
	localAddr = "127.0.0.1:0"
)

var remoteAddr string = config.AppConfig.RemoteMySQL.Host + ":" + config.AppConfig.RemoteMySQL.Port

func SshConnect() (*ssh.Client, error) {
	// SSH 配置
	sshConfig := &ssh.ClientConfig{
		User: config.AppConfig.SSH.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.AppConfig.SSH.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}

	// 建立 SSH 连接
	sshClient, err := ssh.Dial("tcp", config.AppConfig.SSH.Host+":"+config.AppConfig.SSH.Port, sshConfig)
	if err != nil {
		fmt.Printf("SSH 连接失败: %v\n", err)
		return nil, err
	}
	return sshClient, nil
}

// 处理隧道数据转发
func handleTunnel(localConn, remoteConn net.Conn) {
	defer localConn.Close()
	defer remoteConn.Close()

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

// 封装隧道转发逻辑
func StartTunnelForwarding(sshClient *ssh.Client) (int, chan struct{}, error) {
	listener, err := net.Listen("tcp", localAddr)
	if err != nil {
		return 0, nil, fmt.Errorf("本地监听失败: %v", err)
	}

	localPort := listener.Addr().(*net.TCPAddr).Port
	stopChan := make(chan struct{})

	go func() {
		defer listener.Close()
		for {
			select {
			case <-stopChan:
				return
			default:
				localConn, err := listener.Accept()
				if err != nil {
					if !isClosedError(err) {
						fmt.Printf("接受本地连接失败: %v\n", err)
					}
					return
				}

				remoteConn, err := sshClient.Dial("tcp", remoteAddr)
				if err != nil {
					fmt.Printf("SSH 隧道连接 MySQL 失败: %v\n", err)
					localConn.Close()
					return
				}
				go handleTunnel(localConn, remoteConn)
			}
		}
	}()

	return localPort, stopChan, nil
}
func isClosedError(err error) bool {
	if err == nil {
		return false
	}
	if opErr, ok := err.(*net.OpError); ok {
		return opErr.Err.Error() == "use of closed network connection"
	}
	return false
}
