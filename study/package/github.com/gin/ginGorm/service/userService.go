package service

import (
	"study/package/github.com/gin/ginGorm/dao"
	"study/package/github.com/gin/ginGorm/entity"
)

// CreateUser  新建 user信息
func CreateUser(user *entity.User) (err error) {
	if err = dao.Db.Create(user).Error; err != nil {
		return err
	}
	return
}

// GetAllUser 查询所有user
func GetAllUser() (userList []*entity.User, err error) {
	if err = dao.Db.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

// GetUserById 根据id查询user
func GetUserById(id string) (user *entity.User, err error) {
	if err = dao.Db.Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}

// DeleteUserById 根据id删除对应user信息
func DeleteUserById(id string) (err error) {
	err = dao.Db.Where("id=?", id).Delete(&entity.User{}).Error
	return
}

// UpdateUser 更新user信息
func UpdateUser(user entity.User) (err error) {
	err = dao.Db.Save(user).Error
	return
}
