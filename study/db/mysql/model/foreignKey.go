package model

type User struct {
	Uuid  string `gorm:"primaryKey"`
	Name  string
	Age   int
	Class Class `gorm:"foreignKey:UserUuid;references:Uuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
type Class struct {
	Uuid     string `gorm:"primaryKey"`
	UserUuid string `gorm:"column:user_uuid"`
	Name     string
}
