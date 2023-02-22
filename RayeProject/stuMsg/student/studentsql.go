package studentsql

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type myUsualType interface{}

type Student struct {
	ID      int    //20xx44xxx 学号
	Grade   string // 年级  2018级
	Name    string //姓名
	Age     int    //年龄
	Gender  string //性别
	Chinese int    //语文成绩
	Math    int    //数学成绩
	English int    //英语成绩
}

var Db *sql.DB //数据库的连接

func FunctionChoose(button int) (err error) {
	switch button { //button=6的时候进不来，直接退出了
	case 1:
		err = StuDisplay()
		if err != nil {
			return
		}
	case 2:
		var id int
		fmt.Printf("请输入查询的学生学号：")
		_, err = fmt.Scan(&id)
		if err != nil {
			return
		}
		err = StuDisplayOne(id)
		if err != nil {
			return
		}
	case 3:
		var (
			id       int
			item     string
			newValue string //获取的newValue初始都设置成string类型，得到值后，进行类型转换
			a myUsualType
		)
		fmt.Printf("请输入修改信息的学生学号：")
		_, err = fmt.Scan(&id)
		if err != nil {
			return
		}
		//如果学生信息不存在，需要报错返回 查询学生信息的操作
		err = StuDisplayOne(id)
		if err != nil {
			return
		}
		fmt.Printf("请输入修改的信息：（字段 新值）")
		_, err = fmt.Scan(&item, &newValue)
		if err != nil {
			return
		}
		//将获取的item字段更改为小写
		item = strings.ToLower(item)
		//newValue的类型转换
		if item == "id" || item == "age" || item == "chinese" || item == "math" || item == "english" {
			a, err = strconv.ParseInt(newValue, 10, 0) //string 转int64
			if err != nil {
				return
			}
		} else {
			a = fmt.Sprintf("'%v'", newValue)
		}
		err = WriteStu(id, item, a)
		if err != nil {
			return
		}
	case 4:
		var s Student
		fmt.Println(" 请输入新增学生信息：\n", "(学号，年级，姓名，年龄，性别，语文成绩，数学成绩，英语成绩)")
		_, err = fmt.Scan(&s.ID, &s.Grade, &s.Name, &s.Age, &s.Gender, &s.Chinese, &s.Math, &s.English)
		if err != nil {
			return
		}
		err = AddStu(s.ID, s.Grade, s.Name, s.Age, s.Gender, s.Chinese, s.Math, s.English)
		if err != nil {
			return
		}
	case 5:
		var id int
		fmt.Printf("请输入删除的学生学号：")
		_, err = fmt.Scan(&id)
		if err != nil {
			return
		}
		err = DeleteStu(id)
		if err != nil {
			return
		}
	default:
		err = errors.New("输入错误！")
		return
	}
	return err
}

// InitDB 1 数据库的连接 initDB
func InitDB() (err error) {
	//数据库信息
	db := "root:raye12345@tcp(127.0.0.1:3306)/studentMessage"
	//打开数据库
	Db, err = sql.Open("mysql", db)
	if err != nil {
		return
	}
	//尝试与数据库建立连接
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}

// StuDisplayOne 展示当前id的学生信息
func StuDisplayOne(id int) (err error) {
	sqlStr := `select * from studentManage where id=?;`
	row := Db.QueryRow(sqlStr, id)
	var s Student
	err = row.Scan(&s.ID, &s.Grade, &s.Name, &s.Age, &s.Gender, &s.Chinese, &s.Math, &s.English)
	if err != nil {
		err = errors.New("该学号不存在")
	}
	fmt.Printf("%v\n", s)
	return
}

// StuDisplay 学生列表展示 stuDisplay
func StuDisplay() (err error) {
	sqlStr := `select * from studentManage;`
	rows, err := Db.Query(sqlStr)
	if err != nil {
		return
	}
	for rows.Next() {
		var s Student
		_ = rows.Scan(&s.ID, &s.Grade, &s.Name, &s.Age, &s.Gender, &s.Chinese, &s.Math, &s.English)

		fmt.Printf("%#v\n", s)
	}
	return
}

// AddStu addStu 添加学生  插入数据
func AddStu(id int, grade, name string, age int, gender string, chinese, math, english int) (err error) {
	sqlStr := `insert into studentManage(id,grade,name,age,gender,chinese,math,english)values(?,?,?,?,?,?,?,?);`
	_, err = Db.Exec(sqlStr, id, grade, name, age, gender, chinese, math, english)
	if err != nil {
		return
	}
	//展示添加的学生的信息
	fmt.Printf("%#v\n", Student{
		ID:      id,
		Grade:   grade,
		Name:    name,
		Age:     age,
		Gender:  gender,
		Chinese: chinese,
		Math:    math,
		English: english,
	})
	return
}

// WriteStu 编辑学生信息 writeStu 更新数据
func WriteStu(id int, item string, a interface{}) (err error) {
	if item == "id" {
		err = errors.New("学号id不允许更改")
		return
	}
	//sqlStr := `update studentManage set age=? where id=?;`
	sqlStr := fmt.Sprintln("update studentManage set", item, "=", a, " where id=", id, ";") //sql语句，每次编辑一个数据
	//fmt.Printf("%#v\n", sqlStr)
	_, err = Db.Exec(sqlStr)
	if err != nil {
		return
	}
	err = StuDisplayOne(id)
	if err != nil {
		return
	}
	return
}

// DeleteStu 删除学生 deleteStu
func DeleteStu(id int) (err error) {
	sqlStr := `delete from studentManage where id=?;`
	_, err = Db.Exec(sqlStr, id)
	if err != nil {
		return
	}
	err = StuDisplay()
	if err != nil {
		return
	}
	return
}
