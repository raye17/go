package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

type Database struct {
	Driver   string `ini:"driver"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Dbname   string `ini:"dbname"`
}

func main() {
	var database Database
	cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Println("failed to read file :", err)
		os.Exit(1)
	}
	// 典型读取操作，默认分区可以使用空字符串表示
	fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
	fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())
	err = cfg.Section("database").MapTo(&database)
	fmt.Println(database)

}
