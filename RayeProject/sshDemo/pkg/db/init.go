package db

import (
	"fmt"
	"ssh/demo/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbArtwork  *gorm.DB
	dbArtist   *gorm.DB
	dbDigiCopy *gorm.DB
)

const (
	Artwork  = "artwork"
	Artist   = "artist"
	DigiCopy = "digital-copyright"
)

func Init(localPort int) (Db map[string]*gorm.DB, err error) {
	Db = map[string]*gorm.DB{
		Artwork:  dbArtwork,
		Artist:   dbArtist,
		DigiCopy: dbDigiCopy,
	}
	for dbName, _ := range Db {
		sql := config.AppConfig.MySQL[dbName]
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", sql.User, sql.Password, sql.Host, localPort, dbName)
		Db[dbName], err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Printf("打开数据库失败: %v\n", err)
			return
		}
		sqlDB, err1 := Db[dbName].DB()
		if err1 != nil {
			fmt.Printf("获取数据库实例失败: %v\n", err)
			return
		}
		err2 := sqlDB.Ping()
		if err2 != nil {
			fmt.Printf("Ping 数据库失败: %v\n", err)
			return
		}
		fmt.Println("db connected: ", dbName)
		//defer sqlDB.Close()
	}
	return
}
