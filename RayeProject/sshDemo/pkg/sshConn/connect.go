package connect

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"golang.org/x/crypto/ssh"
)

type SSHTunnelService struct {
	sshClient     *ssh.Client
	localListener net.Listener
	stopChan      chan struct{}
	waitGroup     sync.WaitGroup

	sshUser     string
	sshPassword string
	sshServer   string
	localAddr   string
	remoteAddr  string
}

func NewSSHTunnelService(sshUser, sshPassword, sshServer, localAddr, remoteAddr string) *SSHTunnelService {
	return &SSHTunnelService{
		sshUser:     sshUser,
		sshPassword: sshPassword,
		sshServer:   sshServer,
		localAddr:   localAddr,
		remoteAddr:  remoteAddr,
		stopChan:    make(chan struct{}),
	}
}

func (s *SSHTunnelService) StartSSHTunnel() (int, error) {
	// 建立SSH连接
	sshConfig := &ssh.ClientConfig{
		User: s.sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.sshPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	var err error
	s.sshClient, err = ssh.Dial("tcp", s.sshServer, sshConfig)
	if err != nil {
		log.Println("ssh dial failed")
		return 0, err
	}
	s.localListener, err = net.Listen("tcp", s.localAddr)
	if err != nil {
		log.Println("local listen failed")
		return 0, err
	}
	s.waitGroup.Add(1)
	go s.acceptConnections()
	fmt.Println("ssh connect success")
	return s.localListener.Addr().(*net.TCPAddr).Port, nil
}
func (s *SSHTunnelService) acceptConnections() {
	defer s.waitGroup.Done()
	defer s.localListener.Close()
	defer s.sshClient.Close()
	for {
		localConn, err := s.localListener.Accept()
		if err != nil {
			select {
			case <-s.stopChan:
				log.Println("local listener stopped")
				return
			default:
				log.Printf("Error accepting local connection: %v\n", err) // 添加错误日志
			}
			continue
		}
		log.Printf("Accepted local connection from %s", localConn.RemoteAddr())
		s.waitGroup.Add(1)
		go func() {
			defer s.waitGroup.Done()
			defer localConn.Close()
			remoteConn, err := s.sshClient.Dial("tcp", s.remoteAddr)
			if err != nil {
				log.Println("remote dial failed")
				return
			}
			defer remoteConn.Close()
			go copyConn(localConn, remoteConn)
			copyConn(remoteConn, localConn)
			log.Println("remote connect closed")
		}()
	}
}
func copyConn(a, b net.Conn) {
	defer a.Close()
	defer b.Close()
	_, _ = io.Copy(a, b)
}

// Stop 停止 SSH 隧道服务
func (s *SSHTunnelService) Stop() {
	log.Println("Stopping SSH tunnel service...")
	close(s.stopChan)
	if s.localListener != nil {
		s.localListener.Close()
	}
	s.waitGroup.Wait()
	log.Println("SSH tunnel service stopped.")
}
