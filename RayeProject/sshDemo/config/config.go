package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	SSH struct {
		User     string
		Password string
		Host     string
		Port     int
	}
	MySQL struct {
		User     string
		Password string
		Host     string
		Port     int
		Database string
	}
}

var AppConfig Config

func InitConfig() error {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	AppConfig.SSH.User = viper.GetString("ssh.user")
	AppConfig.SSH.Password = viper.GetString("ssh.password")
	AppConfig.SSH.Host = viper.GetString("ssh.host")
	AppConfig.SSH.Port = viper.GetInt("ssh.port")

	AppConfig.MySQL.User = viper.GetString("mysql.user")
	AppConfig.MySQL.Password = viper.GetString("mysql.password")
	AppConfig.MySQL.Host = viper.GetString("mysql.host")
	AppConfig.MySQL.Port = viper.GetInt("mysql.port")
	AppConfig.MySQL.Database = viper.GetString("mysql.database")

	return nil
}
