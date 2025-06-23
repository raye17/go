package main

import (
	"fmt"
	"study/db/common"
	"study/db/common/db"
	"study/db/model"

	_ "github.com/go-sql-driver/mysql"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	db, err := db.DbInit("test03")
	check(err)
	fmt.Println("Successfully connected to the database!")
	err = db.AutoMigrate(model.Student{}, model.Course{})
	if err != nil {
		fmt.Println("Failed to migrate:", err)
	}
	var courses = []*model.Course{
		{
			Id:   1,
			UUID: "7d5ac356-242f-49ad-a80d-a4126f438675",
			Name: "语文",
		},
		{
			Id:   2,
			UUID: "7d5ac356-242f-49ad-a80d-a4126f438675",
			Name: "数学",
		},
		{
			Id:   3,
			UUID: "7d5ac356-242f-49ad-a80d-a4126f438675",
			Name: "英语",
		},
		{
			Id:   4,
			UUID: "7d5ac356-242f-49ad-a80d-a4126f438675",
			Name: "物理",
		},
	}
	// for _, v := range courses {
	// 	uuid, err := common.NewUuid()
	// 	check(err)
	// 	v.UUID = uuid
	// 	fmt.Println(uuid)
	// }
	for _, v := range courses {
		fmt.Println(v)
	}
	uuid, err := common.NewUuid()
	check(err)
	fmt.Println(uuid)
	var student = model.Student{
		Id:     3,
		UUID:   uuid,
		Name:   "wangwu",
		Age:    18,
		Gender: 1,
		//Course: courses,
	}
	if err := db.Create(&student).Error; err != nil {
		fmt.Println(err)
		fmt.Println("create failed")
		return
	}
	fmt.Println("create success")
	// var userInfo model.UserInfo
	// if err := db.Where("id = ?", "1").First(&userInfo).Error; err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("query failed")
	// 	return
	// }
	// var res = userI{
	// 	name: userInfo.Username,
	// 	pad:  userInfo.Password,
	// }
	// var addressList []address
	// if userInfo.ExpressAddress != nil && len(userInfo.ExpressAddress) > 0 {
	// 	for _, v := range userInfo.ExpressAddress {
	// 		addressList = append(addressList, address{
	// 			phone:         v.Phone,
	// 			country:       v.Country,
	// 			province:      v.Province,
	// 			city:          v.City,
	// 			district:      v.District,
	// 			detailAddress: v.DetailAddress,
	// 		})
	// 	}
	// }
	// res.ads = addressList
	// fmt.Println(res)
	// fmt.Println(res.ads[0])
	// fmt.Println(res.ads[1].city)
	// fmt.Printf("%#+v", res)
}
