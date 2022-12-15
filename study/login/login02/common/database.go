package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"loginDemo/model"
)

var DB *gorm.DB

func InitDb() (err error) {
	host := "127.0.0.1"
	port := "3306"
	database := "go_test"
	username := "root"
	password := "raye12345"
	charset := "utf8mb4"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed connect mysql,err:" + err.Error())
	}
	//迁移
	db.AutoMigrate(&model.User{})
	DB = db
	return
}
func GetDb() *gorm.DB {
	return DB
}
