package main

import (
	"fmt"
	"study/package/gin/ginGorm/dao"
	"study/package/gin/ginGorm/entity"
	"study/package/gin/ginGorm/routes"
)

func main() {
	//连接数据库
	err := dao.InitMysql()
	if err != nil {
		fmt.Println(err)
	}
	defer dao.Close()
	//绑定模型
	dao.Db.AutoMigrate(&entity.User{})
	//注册路由
	r := routes.SetRouter()
	r.Run(":8011")
}
