package models

import (
	"gorm.io/gorm"
	"raye/blog/config"
)

// User
type User struct {
	gorm.Model
	Username string `gorm:"column:username;NOT NULL"`
	Password string `gorm:"column:password;NOT NULL"`
	Phone    string `gorm:"column:phone"`
	Email    string `gorm:"column:email"`
	State    int    `gorm:"column:state;default:1;NOT NULL"`
}

// User
func (u *User) User() string {
	return "user"
}

func init() {
	err := config.DB.AutoMigrate(&User{})
	if err != nil {
		return
	}
}
