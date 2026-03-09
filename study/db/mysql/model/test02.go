package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/plugin/soft_delete"
)

type UserInfo struct {
	Id             int32               `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Username       string              `gorm:"column:username;type:varchar(255);" json:"username"`
	Password       string              `gorm:"column:password;type:varchar(255);" json:"password"`
	ExpressAddress ExpressAddressSlice `gorm:"type:json" json:"express_address"`
	CreatedAt      int64               `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      int32               `gorm:"column:updated_at;autoCreateTime"`
	DeletedAt      soft_delete.DeletedAt
}
type ExpressAddress struct {
	Phone         string `gorm:"column:phone;type:varchar(255);" json:"phone"`
	Country       string `gorm:"column:country;type:varchar(255);" json:"country"`
	Province      string `gorm:"column:province;type:varchar(255);" json:"province"`
	City          string `gorm:"column:city;type:varchar(255);" json:"city"`
	District      string `gorm:"column:district;type:varchar(255);" json:"district"`
	DetailAddress string `gorm:"column:detail_address;type:varchar(255);" json:"detail_address"`
}

type ExpressAddressSlice []*ExpressAddress

func (eas ExpressAddressSlice) Value() (driver.Value, error) {
	return json.Marshal(eas)
}

func (eas *ExpressAddressSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, eas)
}
