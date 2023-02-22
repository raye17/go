package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//print系列 将内容输出到系统标准输出
	fmt.Println("*******print系列******")
	fmt.Print("print。 ")
	fmt.Printf("printf:%s\n", "printf")
	fmt.Println("println")
	//output:
	//print。 printf:printf
	//println
	//Fprint系列 会将内容输出到一个io.Writer接口类型的变量w中
	//通常用这个函数往文件中写入内容
	fmt.Println("*******Fprint系列******")
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错,err:", err)
		return
	}
	name := "raye"
	fmt.Fprintf(fileObj, "往文件中写入信息,name:%s", name)
	//Sprint Sprint系列函数会把传入的数据生成并返回一个字符串。
	fmt.Println("*******Sprint系列******")
	s1 := fmt.Sprint("Sprint")
	s2 := fmt.Sprintf("Sprintf:%s", "Sprintf")
	s3 := fmt.Sprintln("Sprintln")
	fmt.Println(s1, s2, s3)
	//Errorf 根据format参数生成格式化字符串并返回一个包含该字符串的错误。
	fmt.Println("*******errorf系列******")
	err = fmt.Errorf("this is a err")
	fmt.Printf("err:%s\ntypeOfErr:%T\n", err, err)
	//bufio.NewReader()
	fmt.Println("*******bufio.NewReader系列******")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please input : ")
	text, err := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}
