package main

import (
	"fmt"
	"study/db/pg/common"
)

func main() {
	db := common.Init()
	err := db.AutoMigrate()
	if err != nil {
		return
	}
	fmt.Println("migrate success")
}
