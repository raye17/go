package controller

import (
	"sxy/demo/pkg/db/model"
	"sxy/demo/pkg/internal/logic"
)

func CreateUser(model model.User) error {
	return logic.CreateUser(model.Username, model.Password)
}

func GetUserInfoByID(id int64) (model.User, error) {
	return logic.GetUserInfoByID(id)
}
func GetUserInfoByUsername(username string) (model.User, error) {
	return logic.GetUserInfoByUsername(username)
}
