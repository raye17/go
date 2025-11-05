package sshCommon

import (
	"context"
	"fmt"
	"net"
	"test001/pkg/service"
	"time"

	"github.com/gin-gonic/gin"
	stdmysql "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type sSHConfig struct {
	SSHUser     string
	SSHPassword string
	SSHHost     string // 格式 ip:port
}

// DBConfig 包含数据库连接信息
type dBConfig struct {
	DBUser string
	DBPass string
	DBHost string // 远程数据库 ip:port
	DBName string
}

var DB *gorm.DB

func UpdateReportList(c *gin.Context) {
	sshCfg := sSHConfig{
		SSHUser:     "root",
		SSHPassword: "",
		SSHHost:     "",
	}
	dbCfg := dBConfig{
		DBUser: "root",
		DBPass: "",
		DBHost: "", // 远程数据库在服务器上的地址
		DBName: "",
	}
	// 1. 建立 SSH 连接
	sshClient, err := ssh.Dial("tcp", sshCfg.SSHHost, &ssh.ClientConfig{
		User: sshCfg.SSHUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshCfg.SSHPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	})
	if err != nil {
		service.Error(c, 201, err, "SSH 连接失败")
	}

	// 2. 注册自定义 DialContext
	protocol := "mysql+ssh"
	stdmysql.RegisterDialContext(protocol, func(ctx context.Context, addr string) (net.Conn, error) {
		return sshClient.Dial("tcp", addr)
	})

	// 3. 构建 DSN
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbCfg.DBUser, dbCfg.DBPass, protocol, dbCfg.DBHost, dbCfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		service.Error(c, 201, err, "数据库连接失败")
	}
	DB = db.Debug()
}
