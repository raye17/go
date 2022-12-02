package login

import (
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	stu "raye.com/stu/student"
)

var (
	registerCodeMode = "raye"
)

// 用户名和密码结构体
type user struct {
	account  string
	password string
}

// FunctionChoose 登录和注册选择
func FunctionChoose(button int) (err error) {
	var (
		account        string //账户名
		password       string //密码
		passwordSecond string //二次确认密码
		registerCode   string //注册码
		ret            bool   //登录返回值
	)
	//注册
	switch button {
	case 0:
		err = login(account, password, ret)
		if err != nil {
			return
		}
	case 1:
		err = register(account, password, passwordSecond, registerCode)
		if err != nil {
			return
		}
		//注册后返回登陆界面
		//延时3s后返回登录界面
		fmt.Println("3s后自动返回登录界面")
		time.Sleep(time.Second * 3)
		timeNow := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(timeNow)
		err = login(account, password, ret)
		if err != nil {
			return
		}
	default:
		err = errors.New("输入错误")
	}
	return
}

// 注册
func register(account, password, passwordSecond, registerCode string) (err error) {
	fmt.Println("请输入注册账号及密码：")
	for {
		_, err = fmt.Scan(&account, &password)
		if err != nil {
			return
		}
		fmt.Println("请确认你的密码：")
		_, err = fmt.Scan(&passwordSecond)
		if err != nil {
			return
		}
		fmt.Println("请输入账号注册码：")
		_, err = fmt.Scan(&registerCode)
		if err != nil {
			return
		}
		err = RegisterAccount(account, password, passwordSecond, registerCode) //注册账号
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
		fmt.Println("请重新输入账户名及密码：")
	}
	return
}

// 登录
func login(account, password string, ret bool) (err error) {
	fmt.Println("请输入账号及密码：")
	for ret != true {
		for i := 0; i < 2; i++ {
			if i == 0 {
				fmt.Print("账号：")
				_, err = fmt.Scan(&account)
				if err != nil {
					return
				}
			}
			if i == 1 {
				fmt.Print("密码：")
				_, err = fmt.Scan(&password)
				if err != nil {
					return
				}
			}
		}
		//2.2.1-----------------验证登录是否成功-----------------
		ret, err = UseLogin(account, password)
		if err != nil {
			fmt.Println(err)
		}
		if ret != true {
			fmt.Println("请重新输入：")
		}

	}
	return
}

// RegisterAccount err=nil则注册成功
func RegisterAccount(account, firstPassword, secondPassword, registerCode string) (err error) {
	if firstPassword != secondPassword {
		err = errors.New("两次输入的密码不一致")
		return
	}
	if registerCode != registerCodeMode {
		err = errors.New("注册码错误")
		return
	}

	sqlQuery := `select account from login where account=?;`
	row := stu.Db.QueryRow(sqlQuery, account)
	var a string
	err = row.Scan(&a)
	if err == nil {
		err = errors.New("账号已存在")
		return
	}
	sqlStr := `insert into login(account,password)values(?,?);`
	_, err = stu.Db.Exec(sqlStr, account, firstPassword)
	if err != nil {
		err = errors.New("注册账号失败")
		return
	}
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("注册成功", timeNow)
	return err
}

// 查询数据库中是否有输入的账户
func accountCheck(account, password string) (err error) {
	//sqlStr := fmt.Sprintf("select * from login where account=%v;", account)
	sqlStr := `select account,password from login where account=?;`
	row := stu.Db.QueryRow(sqlStr, account)
	var u user
	err = row.Scan(&u.account, &u.password)
	if err != nil {
		err = errors.New("账户不存在")
		return
	}
	if u.password != password {
		err = errors.New("密码错误")
		return
	} else {

		return
	}
}

// UseLogin result = 1 时。表示登录成功
func UseLogin(account, password string) (result bool, err error) {
	//查询数据库中是否有account对应的账户
	err = accountCheck(account, password)
	if err != nil {
		return false, err
	}
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("登录成功!", timeNow)
	return true, err
}
