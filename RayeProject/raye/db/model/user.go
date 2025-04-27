package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type User struct {
	ID        uint                  `gorm:"primaryKey;autoIncrement"`
	Name      string                `json:"name"`
	Age       int                   `json:"age"`
	Gender    string                `json:"gender"`
	CreatedAt time.Time             `gorm:"column:created_at;comment:创建时间"`
	UpdatedAt time.Time             `gorm:"column:updated_at;comment:更新时间"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间"`
}
