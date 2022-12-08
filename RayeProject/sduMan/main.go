package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"raye/stu.com/manage"
	"raye/stu.com/menu"
	"raye/stu.com/util"
)

func main() {
	var stuMgs = manage.StudentMgs{
		Students: map[int]manage.Student{},
	}
	for {
		menu.ShowMenu()
		fmt.Print("请输入您的选择：")
		var input int
		_, err := fmt.Scanf("%d\n", &input)
		util.Check(err)
		fmt.Printf("您的选择是:%d\n", input)
		//选项
		switch input {
		case 1:
			manage.GetAll(stuMgs)
		case 2:
			manage.GetOne(stuMgs)
		case 3:
			manage.AddStudent(stuMgs)
		case 4:
			manage.UpdateStudent(stuMgs)
		case 5:
			manage.DeleteStudent(stuMgs)
		case 6:
			os.Exit(0)
		default:
			fmt.Println("invalid input,please enter again")
		}
	}
}
