package model

import "github.com/jinzhu/gorm"

/*
	构建用户结构体User，User继承gorm.Model，可自动添加id等基本字段。
	为了实现收藏与关注功能，每个用户还有一个Collects和Following字段，
	用于保存收藏的文章编号和关注的用户编号，该字段的数据类型为数组
	而UserInfo为部分的用户信息，便于将数据库的查询结果绑定到结构体上

*/

type User struct {
	gorm.Model
	UserName    string `gorm:"varchar(20);not null"`
	PhoneNumber string `gorm:"varchar(20);not null;unique"`
	Password    string `gorm:"size:255;not null"`
	Avatar      string `gorm:"size:255;not null"`
	Collects    Array  `gorm:"type:longtext"`
	Following   Array  `gorm:"type:longtext"`
	Fans        int    `gorm:"AUTO_INCREMENT"`
}
type UserInfo struct {
	ID       uint   `json:"id"`
	Avatar   string `json:"avatar"`
	UserName string `json:"userName"`
}
