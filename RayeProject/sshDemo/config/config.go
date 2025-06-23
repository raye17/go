package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var AppConfig Config

func InitConfig() error {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath("../conf")
	viper.AddConfigPath("../../conf")
	viper.AddConfigPath("../../../conf")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}
	if err := viper.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("解析配置文件失败: %w", err)
	}
	fmt.Println("appconfig:", AppConfig)
	return nil
}

type Config struct {
	System      SystemConfig           `yaml:"system"`
	SSH         SSHConfig              `yaml:"ssh"`
	MySQL       map[string]MySQLConfig `yaml:"mysql"`
	RemoteMySQL RemoteMySQLConfig      `yaml:"remoteMysql"`
}
type SystemConfig struct {
	Mode    string `yaml:"mode"`
	Version string `yaml:"version"`
}
type SSHConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
type RemoteMySQLConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
type MySQLConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Database  string `yaml:"database"`
	Charset   string `yaml:"charset"`
	ParseTime bool   `yaml:"parseTime"`
	Loc       string `yaml:"loc"`
}
