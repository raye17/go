package dao

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

const DRIVER = "mysql"

var Db *gorm.DB

type conf struct {
	Url      string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}

// 获取mysql配置参数
func (c *conf) getConf() *conf {
	//读取config/db.yaml文件
	yamlFile, err := os.ReadFile("./config/db.yaml")
	if err != nil {
		fmt.Println("read file failed:", err)
	}
	//将读取的字符串转换成结构体
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

// InitMysql 初始化连接数据库
func InitMysql() (err error) {
	var c conf
	//获取yaml配置参数
	conf := c.getConf()
	//将yaml文件配置参数拼接成数据库url
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=local",
		conf.UserName,
		conf.Password,
		conf.Url,
		conf.Port,
		conf.Dbname,
	)
	//连接数据库
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("connect mysql failed: ", err)
	}

	//验证数据库是否连接正常
	return
}

// Close 关闭数据库
func Close() {

}
