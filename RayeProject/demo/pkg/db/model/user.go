package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type User struct {
	Id        int64                 `gorm:"primary_key"`
	Username  string                `gorm:"column:username" unique:"username"`
	Password  string                `gorm:"column:password"`
	CreatedAt time.Time             `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time             `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "user"
}
