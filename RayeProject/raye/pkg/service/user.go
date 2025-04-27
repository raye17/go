package service

import (
	"fmt"
	"raye/demo/db"
	"raye/demo/db/model"
)

func CreateUser(user *model.User) error {
	if err := db.DbTest01.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserList() ([]model.User, error) {
	var users []model.User
	if err := db.DbTest01.Find(&users).Error; err != nil {
		return nil, err
	}
	fmt.Println("explain: ", db.DbTest01.Explain(" SELECT * FROM `user` WHERE `user`.`deleted_at` = 0"))
	return users, nil
}

func GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := db.DbTest01.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user *model.User) error {
	if err := db.DbTest01.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uint) error {
	if err := db.DbTest01.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
