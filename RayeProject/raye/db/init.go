package db

import (
	"fmt"
	"raye/demo/config"
	"raye/demo/db/model"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DbTest01 *gorm.DB
var DBtest02 *gorm.DB

func InitDB() (err error) {
	db1 := config.AppConfig.Mysql["test01"]
	db2 := config.AppConfig.Mysql["test02"]
	// 连接数据库
	DbTest01, err = DbConnect(db1)
	if err != nil {
		return err
	}
	DBtest02, err = DbConnect(db2)
	if err != nil {
		return err
	}
	// 自动迁移
	if err := AutoMigrate(); err != nil {
		return err
	}
	return nil
}
func AutoMigrate() error {
	if err := DbTest01.AutoMigrate(model.User{}); err != nil {
		return fmt.Errorf("DbTest01 自动迁移失败: %w", err)
	}
	return nil
}
func DbConnect(cfg config.Mysql) (*gorm.DB, error) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("连接数据库失败: ", err)
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	return db, nil

}
