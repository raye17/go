package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"raye.com/stu/login"
	stu "raye.com/stu/student"
)

/*
 注册需要得到指定的注册码raye才能注册账号
*/

func main() {
	var (
		button1 int
		button2 int
	)
	//1-----------------连接数据库-----------------
	err := stu.InitDB()
	if err != nil {
		return
	}
	//2-----------------登录与注册-----------------
	//2.1 功能选择
	//2.2 注册后登录
	//2.3 登录
	for {
		fmt.Println("请选择登录（输入：0）或注册账号（输入：1）：")
		_, err = fmt.Scan(&button1)
		if err != nil {
			return
		}
		err = login.FunctionChoose(button1)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}
	//3-----------------学生信息管理系统-----------------
	fmt.Println("欢迎访问学生信息管理系统")
	//3.1 功能选择
	for {
		fmt.Println(" 1 查看学生列表\n", "2 查询学生信息\n", "3 修改学生信息\n", "4 增加学生信息\n", "5 删除学生信息\n", "6 退 出 系 统")
		fmt.Printf("请输入：")
		_, err = fmt.Scan(&button2)
		if err != nil {
			return
		}
		if button2 == 6 {
			break
		}
		err = stu.FunctionChoose(button2)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("已退出！")
}
