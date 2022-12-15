package common

//连接数据库
import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	"log"
	"net/url"
	"rayeBlog/model"
)

var (
	Cfg *ini.File
	DB  *gorm.DB
	err error
)

const DRIVER = "mysql"

// InitDB 数据库初始化
func InitDB() *gorm.DB {
	Cfg, err = ini.Load("conf/mysql.ini")
	if err != nil {
		log.Fatalln("failed to parse app.ini: ", err)
	}
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "failed to get section 'database':%v", err)
	}
	dbName := sec.Key("NAME").String()
	user := sec.Key("USER").String()
	password := sec.Key("PASSWORD").String()
	host := sec.Key("HOST").String()
	loc := "Asia/Shanghai"
	db, err := gorm.Open(DRIVER, fmt.Sprintf("%s:%s@tcp(%s)/%s?"+
		"charset=utf8mb4&parseTime=True&loc=%s",
		user,
		password,
		host,
		dbName,
		url.QueryEscape(loc),
	))
	if err != nil {
		panic("failed to open database: " + err.Error())
	}
	//迁移数据表
	db.AutoMigrate(&model.User{})
	DB = db
	return DB

}

// GetDB 数据库信息获取
func GetDB() *gorm.DB {
	return DB
}
