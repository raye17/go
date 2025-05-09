package main

import (
	"study/db/common/db"
	"study/db/model"
)

func main() {
	db, err := db.DbInit("test02")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(model.User{}, model.Class{})
	// userUuid, err := common.NewUuid()
	// if err != nil {
	// 	panic(err)
	// }
	// err = db.Model(model.User{}).Create(&model.User{Uuid: userUuid, Name: "sss", Age: 18}).Error
	// if err != nil {
	// 	panic(err)
	// }
	// classUuid, err := common.NewUuid()
	// if err != nil {
	// 	panic(err)
	// }
	// err = db.Model(model.Class{}).Create(&model.Class{Uuid: classUuid, Name: "大11班", UserUuid: userUuid}).Error
	// if err != nil {
	// 	panic(err)
	// }
	uid := "0ed5f9db-15d9-4e60-831d-5d7892fc2c4c"
	err = db.Model(&model.User{}).Where("uuid = ?", uid).Delete(&model.User{}).Error
	if err != nil {
		panic(err)
	}

}
