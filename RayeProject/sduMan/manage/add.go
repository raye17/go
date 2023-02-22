package manage

import (
	"fmt"
	"raye/stu.com/util"
)

func (stuMgs StudentMgs) addStudent(s Student) {
	stuMgs.Students[s.id] = s
}
func AddStudent(stuMgs StudentMgs) {
	var (
		id     int
		name   string
		age    int
		gender string
		class  string
	)
againInput:
	fmt.Println("请输入学生学号：")
	_, err := fmt.Scanln(&id)
	util.Check(err)
	if id == 0 {
		fmt.Println("学号不能为0，请重新输入：")
		goto againInput
	}
	res := stuMgs.isExist(id)
	if res {
		fmt.Println("该学生已存在！")
		return
	}
	for i := 0; i < 4; i++ {
		if i == 0 {
			fmt.Println("请输入学生姓名：")
			_, err := fmt.Scanln(&name)
			util.Check(err)
		}
		if i == 1 {
			fmt.Println("请输入学生年龄：")
			_, err := fmt.Scanln(&age)
			util.Check(err)
		}
		if i == 2 {
			fmt.Println("请输入学生性别：")
			_, err := fmt.Scanln(&gender)
			util.Check(err)
		}
		if i == 3 {
			fmt.Println("请输入学生班级：")
			_, err := fmt.Scanln(&class)
			util.Check(err)
		}

	}
	stuInfo := Student{
		id:     id,
		name:   name,
		age:    age,
		gender: gender,
		class:  class,
	}
	stuMgs.addStudent(stuInfo)
	util.EditFile(stuInfo)
	fmt.Println("add student information success!")

}
