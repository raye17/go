package main

import "fmt"

// 占位符
type human struct {
	name string
	age  int
}

func main() {
	var people = human{name: "raye", age: 22}
	//普通占位符
	fmt.Printf("%v   ", people)  //相应值得默认格式
	fmt.Printf("%+v   ", people) //打印结构体时会添加字段名
	fmt.Printf("%#v   ", people) //相应值的go语法表示
	fmt.Printf("%T   ", people)  //相应值类型的go语法表示
	fmt.Printf("%%\n")           //字面上的百分号
	//布尔占位符
	fmt.Printf("%t,%t\n", true, false) //true or false
	//整数占位符
	fmt.Printf("%b  ", 5)      //二进制
	fmt.Printf("%c  ", 0x4E2D) //Unicode码所表示字符
	fmt.Printf("%d  ", 19)     //十进制
	fmt.Printf("%o  ", 10)     //八进制
	fmt.Printf("%q  ", 0x4E2D) //单引号围绕
	fmt.Printf("%x  ", 10)     //十六进制，小写字母表示
	fmt.Printf("%X  ", 100)    //十六进制，字母大写
	fmt.Printf("%U", 0x4E2D)   //Unicode 格式
}
