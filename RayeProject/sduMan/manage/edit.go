package manage

import "fmt"

func (stuMgs StudentMgs) updateStudent(s Student) {
	stuMgs.Students[s.id] = s
}
func UpdateStudent(stuMgs StudentMgs) {
	var id int
	fmt.Println("请输入要修改信息的学生学号：")
	fmt.Scanln(&id)
	stuInfo := stuMgs.getStudent(id)
	fmt.Println("当前学生信息是=>学号：", stuInfo.id, "姓名",
		stuInfo.name, "年龄", stuInfo.age, "性别：", stuInfo.gender, "班级", stuInfo.class)
}
