package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"loginDemo/model"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "go_test"
	username := "root"
	password := "raye12345"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed connect mysql,err:" + err.Error())
	}
	//迁移
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}
func GetDb() *gorm.DB {
	return DB
}
