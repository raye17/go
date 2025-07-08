package db

import (
	"fmt"
	"sxy/demo/config"
	"sxy/demo/pkg/db/model"
	zaplog "sxy/demo/pkg/zap"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func DbInit() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		config.AppConfig.Mysql.User,
		config.AppConfig.Mysql.Password,
		config.AppConfig.Mysql.Host,
		config.AppConfig.Mysql.Port,
		config.AppConfig.Mysql.DBName,
		config.AppConfig.Mysql.Charset,
		config.AppConfig.Mysql.ParseTime,
		config.AppConfig.Mysql.Loc,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		zaplog.Logger.Error("gorm.Open() failed, err:", zap.Error(err))
		return
	}
	DB = db
	if err := migrate(); err != nil {
		zaplog.Logger.Error("migrate() failed, err:", zap.Error(err))
		return
	}
	zaplog.Logger.Info("db init success")
}
func migrate() error {
	return DB.AutoMigrate(&model.User{})
}
