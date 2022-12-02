package manage

type Student struct {
	id     int //学号
	name   string
	age    int
	gender string
	class  string
}

// StudentMgs 学生信息管理
// 定义管理者结构体，管理者中存储着学生信息
type StudentMgs struct {
	Students map[int]Student
}
