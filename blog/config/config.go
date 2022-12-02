package config

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"os"
	"unicode/utf8"
)

type configData struct {
	DB     mysqlConfig  `json:"mysql"`
	Server serverConfig `json:"server"`
}

/*
上面主要是获取运行环境的目录，来确定项目目录，它有2种处理方法，
一种是使用执行文件所在的目录，另一种是使用执行命令时所在的目录
*/
//定义执行目录
func initPath() {
	sep := string(os.PathSeparator)
	//执行文件所在目录
	//root := filepath.Dir(os.Args[0])
	//ExecPath, _ := filepath.Abs(root)
	//执行命令所在目录
	var err error
	ExecPath, err = os.Getwd()
	if err != nil {
		fmt.Println("error:", err)
	}
	length := utf8.RuneCountInString(ExecPath)
	lastChar := ExecPath[length-1:]
	if lastChar != sep {
		ExecPath = ExecPath + sep
	}
}

// InitJSON 读取json文件
func InitJSON() {
	rawConfig, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		fmt.Println("Invalid Config:", err.Error())
		os.Exit(-1)
	}
	if err := json.Unmarshal(rawConfig, &JsonData); err != nil {
		fmt.Println("Invalid Config:", err.Error())
		os.Exit(-1)
	}
}

// 解析server
func initServer() {
	ServerConfig = JsonData.Server
}

// InitDb 解析mysql
func InitDb(setting *mysqlConfig) error {
	var (
		db  *gorm.DB
		err error
	)
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		setting.User, setting.Password, setting.Host, setting.Port, setting.Database)
	setting.Url = url
	db, err = gorm.Open(mysql.Open(url), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //禁用表名加s
		},
		Logger:                                   logger.Default.LogMode(logger.Info), // 打印sql语句
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用创建外键约束
	})
	if err != nil {
		panic("Connecting database failed: " + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(1000)
	sqlDB.SetMaxOpenConns(100000)
	sqlDB.SetConnMaxLifetime(-1)

	DB = db

	return nil
}

var ExecPath string
var JsonData configData
var ServerConfig serverConfig
var DB *gorm.DB

func init() {
	initPath()
	//读取json
	InitJSON()
	//读取server
	initServer()
	//初始化数据库
	err := InitDb(&JsonData.DB)
	if err != nil {
		fmt.Println("Failed To Connect Database: ", err.Error())
		os.Exit(-1)
	}
}
func GetDB() *gorm.DB {
	return DB
}
