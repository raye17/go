package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var AppConfig Config

func InitConfig() error {

	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf/local")
	viper.AddConfigPath("../conf")       // 上级目录下的conf
	viper.AddConfigPath("../../conf")    // 上两级目录下的conf
	viper.AddConfigPath("../../../conf") // 上三级目录下的conf
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("解析配置文件失败: %w", err)
	}

	return nil
}

type Config struct {
	System struct {
		Mode    string `yaml:"mode"`
		Version string `yaml:"version"`
	}
	Mysql map[string]Mysql
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		Database int    `yaml:"database"`
	}
	SMTP     SMTPConfig
	RabbitMQ RabbitMQConfig
	Oss      OssConfig
}
type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type RabbitMQConfig struct {
	URL string
}
type OssConfig struct {
	EndPoint        string
	AccessKeyId     string
	AccessKeySecret string
	Bucket          string
	Region          string
}
