package common

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	fmt.Println("pg init")
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "123456"
	dbname := "szai"
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(conn)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}
	autoMigrate(db)
	return db
}
func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate()
}
