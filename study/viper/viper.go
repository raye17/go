package viper

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func LoadConfig() {
	fmt.Println(os.Getwd())
	//viper.SetConfigName("test")
	//viper.SetConfigType("yaml")
	viper.SetConfigFile("./env/local/test.yaml")
	//viper.AddConfigPath("./env/local")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("mysql: ", viper.Get("mysql"))
	fmt.Println("redis: ", viper.Get("redis"))
}
