package model

import (
	"gorm.io/plugin/soft_delete"
)

type Student struct {
	Id        int32     `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	UUID      string    `json:"uuid" gorm:"column:uuid;type:varchar(1024);comment:UUID"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(2048);comment:学生姓名"`
	Age       int32     `json:"age" gorm:"column:age;type:int(11);comment:年龄"`
	Gender    int32     `json:"gender" gorm:"column:gender;type:int(11);comment:性别 1:男 2:女"`
	Course    []*Course `gorm:"foreignKey:UUID;references:UUID"`
	CreatedAt int64     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt int64     `gorm:"column:updated_at;autoCreateTime"`
	DeletedAt soft_delete.DeletedAt
}
type Course struct {
	Id        int32  `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	UUID      string `json:"uuid" gorm:"column:uuid;type:varchar(1024);comment:UUID"`
	Name      string `json:"name" gorm:"column:name;type:varchar(2048);comment:课程名称"`
	Teacher   string `json:"teacher" gorm:"column:teacher;type:varchar(2048);comment:任课老师"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime"`
	DeletedAt soft_delete.DeletedAt
}
