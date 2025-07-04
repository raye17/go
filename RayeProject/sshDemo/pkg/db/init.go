package db

import (
	"fmt"
	"ssh/demo/config"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB map[string]*gorm.DB

func DBInit() {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	DB = make(map[string]*gorm.DB)
	for dbName, dbConfig := range config.AppConfig.MySQL {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Database,
		)
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       dsn,   // DSN data source name
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据版本自动配置
		}), &gorm.Config{
			Logger: ormLogger,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			panic(err)
		}
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(20)  //设置连接池，空闲
		sqlDB.SetMaxOpenConns(100) //打开
		sqlDB.SetConnMaxLifetime(time.Second * 30)
		// if err = db.AutoMigrate(); err != nil {
		// 	panic(err)
		// }
		DB[dbName] = db
	}
}
