package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var AppConfig Config

func InitConfig() error {
	if wd, err := os.Getwd(); err == nil {
		fmt.Println("当前工作目录:", wd)
	}
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
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
}
type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
