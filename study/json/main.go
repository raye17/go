package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string
	Id     int
	Age    interface{}
	Gender string
}
type StudentInfo struct {
	Class string
	Score int
}

func main() {
	student := Student{
		Name:   "raye",
		Id:     36001,
		Age:    18,
		Gender: "男",
	}
	data, err := json.Marshal(&student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", string(data))
	fmt.Println()
	fmt.Println("*********")
	str := "{\"Class\":\"36班\",\"Score\":100}"
	var stu StudentInfo
	err = json.Unmarshal([]byte(str), &stu)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", stu)
	fmt.Println(stu.Class, stu.Score)
}
