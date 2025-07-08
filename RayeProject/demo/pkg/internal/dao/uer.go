package dao

import (
	"sxy/demo/pkg/db"
	"sxy/demo/pkg/db/model"
)

func CreateUser(user *model.User) error {
	err := db.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateUser(user *model.User) error {
	err := db.DB.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}
func DeleteUser(id int64) error {
	err := db.DB.Delete(&model.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
func GetUserInfoByID(id int64, user *model.User) error {
	err := db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func GetUserInfoByUsername(username string, user *model.User) error {
	err := db.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}
