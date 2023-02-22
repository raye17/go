package manage

import (
	"fmt"
	"raye/stu.com/util"
)

func (stuMgs StudentMgs) deleteStudent(id int) bool {
	_, exist := stuMgs.Students[id]
	if !exist {
		return false
	}
	delete(stuMgs.Students, id)
	return true
}
func DeleteStudent(stuMgs StudentMgs) {
	var id int
	fmt.Println("请输入要删除的学生学号：")
	_, err := fmt.Scanln(&id)
	util.Check(err)
	res := stuMgs.deleteStudent(id)
	if res {
		fmt.Println("删除成功！")
	} else {
		fmt.Println("该学生不存在！")
	}
}
