package main

import (
	"fmt"
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
	db, err := db.DbInit("yourapi")
	check(err)
	fmt.Println("Successfully connected to the database!")
	err = db.AutoMigrate(model.UserInfo{})
	if err != nil {
		fmt.Println("Failed to migrate:", err)
	}
	//expressAddressList := []*model.ExpressAddress{}
	// expressAddressList = append(expressAddressList, &model.ExpressAddress{
	// 	Phone:         string("11111111111"),
	// 	Province:      string("北京市"),
	// 	City:          string("北京市"),
	// 	District:      string("海淀区"),
	// 	DetailAddress: string("中关村大街"),
	// })
	// expressAddressList = append(expressAddressList, &model.ExpressAddress{
	// 	Phone:         string("222222222222"),
	// 	Province:      string("北京市"),
	// 	City:          string("北京市"),
	// 	District:      string("海淀区"),
	// 	DetailAddress: string("中关村大街"),
	// })
	// userInfo := model.UserInfo{
	// 	Username:       string("sxy"),
	// 	Password:       string("123456"),
	// 	ExpressAddress: expressAddressList,
	// }
	// if err := db.Create(&userInfo).Error; err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("create failed")
	// 	return
	// }
	//fmt.Println("create success")
	var userInfo model.UserInfo
	if err := db.Where("id = ?", "1").First(&userInfo).Error; err != nil {
		fmt.Println(err)
		fmt.Println("query failed")
		return
	}
	var res = userI{
		name: userInfo.Username,
		pad:  userInfo.Password,
	}
	var addressList []address
	if userInfo.ExpressAddress != nil && len(userInfo.ExpressAddress) > 0 {
		for _, v := range userInfo.ExpressAddress {
			addressList = append(addressList, address{
				phone:         v.Phone,
				country:       v.Country,
				province:      v.Province,
				city:          v.City,
				district:      v.District,
				detailAddress: v.DetailAddress,
			})
		}
	}
	res.ads = addressList
	fmt.Println(res)
	fmt.Println(res.ads[0])
	fmt.Println(res.ads[1].city)
	fmt.Printf("%#+v", res)
}

type userI struct {
	name string
	pad  string
	ads  []address
}
type address struct {
	phone         string
	country       string
	province      string
	city          string
	district      string
	detailAddress string
}
