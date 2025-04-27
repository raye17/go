package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"ssh/demo/pkg/service"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"golang.org/x/crypto/ssh"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

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

// 检查是否为连接关闭错误
func isClosedError(err error) bool {
	if err == nil {
		return false
	}
	if opErr, ok := err.(*net.OpError); ok {
		return opErr.Err.Error() == "use of closed network connection"
	}
	return false
}
func sshConnect() (*ssh.Client, error) {
	// SSH 配置
	sshConfig := &ssh.ClientConfig{
		User: "",
		Auth: []ssh.AuthMethod{
			ssh.Password(""),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}

	// 建立 SSH 连接
	sshClient, err := ssh.Dial("tcp", "", sshConfig)
	if err != nil {
		fmt.Printf("SSH 连接失败: %v\n", err)
		return nil, err
	}
	return sshClient, nil
}

// 封装隧道转发逻辑
func startTunnelForwarding(sshClient *ssh.Client, localAddr, remoteAddr string) (int, chan struct{}, error) {
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

type ResultRecord struct {
	UUID    string
	Message string
}

func main() {
	service.Ser()
	sshClient, err := sshConnect()
	if err != nil {
		return
	}
	defer sshClient.Close()

	// 使用封装后的函数
	localPort, stopChan, err := startTunnelForwarding(
		sshClient,
		"127.0.0.1:0",
		"",
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("本地隧道端口: %d\n", localPort)

	// 等待隧道准备就绪
	time.Sleep(1 * time.Second)

	// MySQL 连接
	dsn1 := fmt.Sprintf("fonchain_sunxiaoyang:cads234jpaGJ@va21@tcp(127.0.0.1:%d)/artwork?charset=utf8mb4&parseTime=True&loc=Local", localPort)
	dsn2 := fmt.Sprintf("fonchain_sunxiaoyang:cads234jpaGJ@va21@tcp(127.0.0.1:%d)/artist?charset=utf8mb4&parseTime=True&loc=Local", localPort)
	dsn3 := fmt.Sprintf("fonchain_sunxiaoyang:cads234jpaGJ@va21@tcp(127.0.0.1:%d)/digital-copyright?charset=utf8mb4&parseTime=True&loc=Local", localPort)
	DbArtwork, err := gorm.Open(mysql.Open(dsn1), &gorm.Config{})
	if err != nil {
		fmt.Printf("打开数据库失败: %v\n", err)
		return
	}
	DbArtist, err := gorm.Open(mysql.Open(dsn2), &gorm.Config{})
	if err != nil {
		fmt.Printf("打开数据库失败: %v\n", err)
		return
	}
	DbDigiCopy, err := gorm.Open(mysql.Open(dsn3), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("打开digital数据库失败: %v\n", err)
		return
	}
	// 测试连接
	sqlDB, err := DbArtwork.DB()
	if err != nil {
		fmt.Printf("获取数据库实例失败: %v\n", err)
		return
	}
	err = sqlDB.Ping()
	if err != nil {
		fmt.Printf("Ping 数据库失败: %v\n", err)
		return
	}
	artistDb, err := DbArtist.DB()
	if err != nil {
		fmt.Printf("获取artist数据库实例失败: %v\n", err)
		return
	}
	err = artistDb.Ping()
	if err != nil {
		fmt.Printf("Ping artist数据库失败: %v\n", err)
		return
	}
	digiDb, err := DbDigiCopy.DB()
	if err != nil {
		fmt.Printf("获取di数据库实例失败: %v\n", err)
		return
	}
	err = digiDb.Ping()
	if err != nil {
		fmt.Printf("Ping di数据库失败: %v\n", err)
		return
	}
	defer sqlDB.Close()
	defer artistDb.Close()
	defer digiDb.Close()

	fmt.Println("Successfully connected to MySQL via SSH!")

	// 停止隧道
	close(stopChan)
}

// 修改文件写入部分
func writeRecordsToFile(records []ResultRecord, filename string) error {
	var lines []string
	for _, r := range records {
		lines = append(lines, fmt.Sprintf("%s\t%s", r.UUID, r.Message))
	}

	// 检查文件是否存在
	if _, err := os.Stat(filename); err == nil {
		// 文件存在，以追加模式打开
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		// 写入换行符和新增内容
		if _, err = file.WriteString("\n" + strings.Join(lines, "\n")); err != nil {
			return err
		}
	} else {
		// 文件不存在，创建新文件
		return os.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
	}
	return nil
}
func NewSf() *snowflake.Node {
	var err error
	var st time.Time
	nodeNum, _ := strconv.Atoi("5")
	st, err = time.Parse("2006-01-02", "2023-05-31")
	if err != nil {
		panic(err)
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, errS := snowflake.NewNode(int64(nodeNum))
	if errS != nil {
		panic(errS)
	}
	return node
}
