package db

import (
	dciConfig "chain-dci/config"
	"gorm.io/gorm"
	"strings"
)

//var DciProvider = wire.NewSet(NewDci)

func NewDci() *gorm.DB {
	connDci := strings.Join([]string{dciConfig.Data.ChainDci.User, ":", dciConfig.Data.ChainDci.Password,
		"@tcp(", dciConfig.Data.ChainDci.Host, ":", dciConfig.Data.ChainDci.Port, ")/",
		dciConfig.Data.ChainDci.DbName, "?charset=utf8mb4&parseTime=true"}, "")
	DciDB := loadMysqlConn(connDci)
	return DciDB
}
