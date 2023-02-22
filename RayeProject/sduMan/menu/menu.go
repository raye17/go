package menu

//展示菜单
import "fmt"

func ShowMenu() {
	fmt.Println(`
	welcome to student-manange system:
		 "1 展示所有学员信息"
		 "2 查询单个学员信息"
 		 "3 增加学员"
		 "4 编辑学员信息"
		 "5 删除学员" 
		 "6 退出"
	`)
}
