package db

import (
	"database/sql"
	"fmt"
	"net"

	"ssh/demo/config"
	"ssh/demo/db/ssh"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() (*sql.DB, net.Listener, error) {
	// 建立SSH隧道连接
	listener, err := ssh.GetTunnelConnection()
	if err != nil {
		return nil, nil, fmt.Errorf("SSH隧道连接失败: %w", err)
	}
	localPort := listener.Addr().(*net.TCPAddr).Port
	fmt.Printf("本地隧道端口: %d\n", localPort)
	// 初始化数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConfig.MySQL.User,
		config.AppConfig.MySQL.Password,
		localPort,
		config.AppConfig.MySQL.Database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("打开数据库失败: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, nil, fmt.Errorf("Ping 数据库失败: %w", err)
	}
	DB = db
	return db, listener, nil
}
