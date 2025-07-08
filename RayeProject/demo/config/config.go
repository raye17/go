package config

import (
	zaplog "sxy/demo/pkg/zap"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var AppConfig *Config

type Config struct {
	System System
	Mysql  MySQL
	Redis  Redis
}
type System struct {
	Env      string
	HttpPort int
}
type MySQL struct {
	Host      string
	Port      string
	User      string
	Password  string
	DBName    string
	Charset   string
	ParseTime string
	Loc       string
}
type Redis struct {
	Addr     string
	Password string
	DB       int
}

func InitConfig() error {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../conf")
	err := viper.ReadInConfig()
	if err != nil {
		zaplog.Logger.Error("read config failed, err:", zap.Error(err))
		return err
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		zaplog.Logger.Error("unmarshal config failed, err:", zap.Error(err))
		return err
	}
	zaplog.Logger.Info("config init success")
	return nil
}
