package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"ssh/demo/config"
	connect "ssh/demo/pkg/sshConn"
	"syscall"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	failedRecoed  = "../data/failed_uuids-5-23.txt"
	successRecoed = "../data/success_uuids-5-23.txt"
)

func init() {
	err := config.InitConfig()
	if err != nil {
		zap.L().Error("config.InitConfig() failed", zap.Error(err))
		return
	}
	//db.DBInit()
}
func main() {
	//sshProd.SshConnect()
	var result map[string]any
	//db.DB["test01"].Table("user").Find(&result)
	//fmt.Println(result)
	//res := ReturnMap()
	//fmt.Println(res["sss"])
	sshService := connect.NewSSHTunnelService(
		config.AppConfig.SSH.User,
		config.AppConfig.SSH.Password,
		config.AppConfig.SSH.Host+":"+config.AppConfig.SSH.Port,
		"127.0.0.1:0",
		config.AppConfig.RemoteMySQL.Host+":"+config.AppConfig.RemoteMySQL.Port,
	)
	localPort, err := sshService.StartSSHTunnel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(localPort)
	dbName := "artwork"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConfig.MySQL[dbName].User,
		config.AppConfig.MySQL[dbName].Password,
		localPort,
		dbName,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		//Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("err: ", err)
		panic(err)
	}
	db.Table("artwork_profile").Where("tfnum = ? ", "T08621010").Find(&result)
	fmt.Println(result)
	// 阻塞主 Goroutine，等待中断信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	sshService.Stop()
	log.Println("Application exited.")
}
