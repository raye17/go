package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg          *ini.File
	RunMode      string
	HttpPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
	PageSize     int
	JwtSecret    string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalln("failed to parse app.ini: ", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalln("failed to get section 'server':", err)
	}
	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeOut = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeOut = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalln("failed to get section 'app':", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
