package service

import (
	"raye/demo/db"
	"raye/demo/db/model"
)

func CreateUser(user *model.User) error {
	if err := db.DbTest01.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
