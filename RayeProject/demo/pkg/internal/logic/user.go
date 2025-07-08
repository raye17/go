package logic

import (
	"errors"
	"sxy/demo/pkg/db/model"
	"sxy/demo/pkg/internal/dao"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password string) error {
	user := &model.User{
		Username: username,
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	err = dao.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
func GetUserInfoByID(id int64) (model.User, error) {
	var user model.User
	if id == 0 {
		return user, errors.New("id is empty")
	}
	if err := dao.GetUserInfoByID(id, &user); err != nil {
		return user, err
	}
	return user, nil
}
func GetUserInfoByUsername(username string) (model.User, error) {
	var user model.User
	if username == "" {
		return user, errors.New("username is empty")
	}
	if err := dao.GetUserInfoByUsername(username, &user); err != nil {
		return user, err
	}
	return user, nil
}
