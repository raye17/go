package service

import (
	"fmt"
	"raye/demo/db"
	"raye/demo/db/model"
	"raye/demo/pkg/mq"
	"raye/demo/pkg/utlis/e"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type user struct {
	Id       uint
	Name     string
	Password string
	Age      int
	Gender   string
}

func CreateUser(c *gin.Context) {
	var req user
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("bind json error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
	}
	if err := db.DbTest01.Create(&req).Error; err != nil {
		fmt.Println("create user error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
	}
	ResponseMsg(c, e.Success, e.SuccessMsg, nil, "ok")
}

func GetUserList(c *gin.Context) {
	var users []model.User
	if err := db.DbTest01.Find(&users).Error; err != nil {
		fmt.Println("get user list error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
	}
	fmt.Println("explain: ", db.DbTest01.Explain(" SELECT * FROM `user` WHERE `user`.`deleted_at` = 0"))
	ResponseMsg(c, e.Success, e.SuccessMsg, nil, users)
}

func GetUserByID(c *gin.Context) {
	var req user
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("bind json error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
	}
	var user model.User
	if err := db.DbTest01.First(&user, req).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			ResponseMsg(c, e.Success, "user not found", nil, nil)
			return
		}
		fmt.Println("get user by id error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
	}
	go func() {
		_ = mq.PushMsg(user)
	}()
	ResponseMsg(c, e.Success, e.SuccessMsg, nil, user)
}

func UpdateUser(c *gin.Context) {
	var req user
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("bind json error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
	}
	var user model.User
	if err := db.DbTest01.Save(&user).Error; err != nil {
		fmt.Println("update user error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
	}
	ResponseMsg(c, e.Success, e.SuccessMsg, nil, user)
}

func DeleteUser(c *gin.Context) {
	var req user
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("bind json error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
	}
	if err := db.DbTest01.Delete(&model.User{}, req).Error; err != nil {
		fmt.Println("delete user error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
	}
	ResponseMsg(c, e.Success, e.SuccessMsg, nil, nil)
}
