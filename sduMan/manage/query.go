package manage

import "fmt"

func (stuMgs StudentMgs) listStudent() map[int]Student {
	return stuMgs.Students
}
func (stuMgs StudentMgs) getStudent(id int) Student {
	return stuMgs.Students[id]
}
func GetAll(stuMgs StudentMgs) {
	fmt.Println("以下是所有学生信息")
	fmt.Println("************Start***************")
	students := stuMgs.listStudent()
	for _, stuInfo := range students {
		fmt.Println("学号：", stuInfo.id, "姓名：", stuInfo.name, "年龄：", stuInfo.age,
			"性别：", stuInfo.gender, "班级：", stuInfo.class)
	}
	fmt.Println("************END*****************")

}
func GetOne(stuMgs StudentMgs) {
	fmt.Print("请输入要查询学生学号：")
	var id int
	fmt.Scanln(&id)
	stuInfo := stuMgs.getStudent(id)
	if stuInfo.id == 0 {
		fmt.Println("该学生不存在!")
	} else {
		fmt.Println("************Start***************")
		fmt.Println("学号：", stuInfo.id, "姓名：", stuInfo.name, "年龄：", stuInfo.age,
			"性别：", stuInfo.gender, "班级：", stuInfo.class)
		fmt.Println("************END****************")
	}
}
