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

func InitDB() error {
	db1 := config.AppConfig.Mysql["test01"]
	db2 := config.AppConfig.Mysql["test02"]
	// 连接数据库
	db1Conn, err := DbConnect(db1)
	if err != nil {
		return err
	}
	db2Conn, err := DbConnect(db2)
	if err != nil {
		return err
	}
	DbTest01 = db1Conn
	DBtest02 = db2Conn
	// 自动迁移
	if err := AutoMigrateForDbTest01(model.User{}); err != nil {
		return err
	}
	return nil

}

// 初始化数据库表
func AutoMigrateForDbTest01(models ...interface{}) error {
	if err := DbTest01.AutoMigrate(models...); err != nil {
		return fmt.Errorf("DbTest01 自动迁移失败: %w", err)
	}
	return nil
}

func AutoMigrateForDBtest02(models ...interface{}) error {
	if err := DBtest02.AutoMigrate(models...); err != nil {
		return fmt.Errorf("DBtest02 自动迁移失败: %w", err)
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
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	return db, nil

}
