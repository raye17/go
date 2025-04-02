package main

import (
	"fmt"
	"ssh/demo/config"
	"ssh/demo/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		fmt.Printf("读取文件配置错误: %v\n", err)
		return
	}
	// 初始化配置并建立连接
	db, listener, err := db.Init()
	if err != nil {
		fmt.Printf("初始化失败: %v\n", err)
		return
	}
	defer listener.Close()
	defer db.Close()

	stopChan := make(chan struct{})

	fmt.Println("Successfully connected to MySQL via SSH!")
	// 导出user表数据到CSV
	// err = service.ExportUserToCSV(db)
	// if err != nil {
	// 	fmt.Printf("导出user表数据失败: %v\n", err)
	// 	return
	// }
	close(stopChan)

}
