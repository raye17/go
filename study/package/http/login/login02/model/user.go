package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"varchar(20);not null"`
	Password string `gorm:"size:255;not null"`
}
