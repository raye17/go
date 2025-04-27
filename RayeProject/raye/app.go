package main

import (
	"context"
	"encoding/json"
	"fmt"
	"raye/demo/config"
	"raye/demo/db"
	"raye/demo/pkg/service"
	"raye/demo/redis/cache"
	"time"

	"gorm.io/gorm"
)

type userInfo struct {
	Id        uint
	Name      string
	Age       int
	Gender    string
	CreatedAt string
	UpdatedAt string
}

func main() {
	err := config.InitConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config.AppConfig)
	err = db.InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	list, err := service.GetUserList()
	if err != nil {
		fmt.Println(err)
		return
	}
	var userList []userInfo
	for _, v := range list {
		userList = append(userList, userInfo{
			Id:        v.ID,
			Name:      v.Name,
			Age:       v.Age,
			Gender:    v.Gender,
			CreatedAt: v.CreatedAt.Format("2006-01-02"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	for _, v := range userList {
		fmt.Println(v)
		userJSON, err := json.Marshal(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = cache.RedisClient.Set(context.Background(), v.Name, userJSON, 10*time.Hour).Result()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	var user userInfo
	users, err := service.GetUserByID(3)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			fmt.Println("用户不存在")
		} else {
			fmt.Println(err)
			return
		}
	} else {
		user.Id = users.ID
		user.Name = users.Name
		user.Age = users.Age
		user.Gender = users.Gender
		user.CreatedAt = users.CreatedAt.Format("2006-01-02 15:04:05")
		user.UpdatedAt = users.UpdatedAt.Format("2006-01-02 15:04:05")
		fmt.Println(user)
	}
	var u userInfo
	res, err := cache.RedisClient.Get(context.Background(), "sxy").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	err = json.Unmarshal([]byte(res), &u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("json unmarshal: ")
	fmt.Println(u)
	fmt.Println("main over")
}
func (u userInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
