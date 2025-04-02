package ssh

import (
	"fmt"
	"ssh/demo/config"
	"time"

	"golang.org/x/crypto/ssh"
)

// GetSSHClient 建立并返回SSH客户端连接
func GetSSHClient() (*ssh.Client, error) {
	sshConfig := &ssh.ClientConfig{
		User: config.AppConfig.SSH.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.AppConfig.SSH.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}

	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", config.AppConfig.SSH.Host, config.AppConfig.SSH.Port), sshConfig)
	if err != nil {
		return nil, fmt.Errorf("SSH连接失败: %w", err)
	}

	return sshClient, nil
}
