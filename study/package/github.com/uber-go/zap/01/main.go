package main

import (
	"fmt"
	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	logger, _ = zap.NewProduction()
	//log1 := zap.NewExample()
	//log2, _ := zap.NewDevelopment()
	//defer logger.Sync()
	//logger.Info("001-info")
	//ssl := logger.Named("ssl")
	//ssl.Info("002-info ")
	//ssl.Named("001").Info("003-info")
	//ssl.Named("888").Info("004-info")
	//ssl.Warn("warn")
	//fmt.Println(logger.Level())
	//fmt.Println(log1.Level())
	//fmt.Println(log2.Level())
	l := zap.L()
	l.Info("lll")
	fmt.Println(l.Level())

}
