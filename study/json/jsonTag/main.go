package main

import (
	"encoding/json"
	"fmt"
)

// Student Tag：-表示不进行序列化
// omitempty 序列化的时候忽略零值或空值
type Student struct {
	Name   string `json:"name"`
	Id     int    `json:"-"`
	Age    int    `json:"age,omitempty"`
	Sage   int    `json:"sage"`
	Gender string `json:"gender"`
}

func main() {
	stu := &Student{
		Name:   "raye",
		Id:     36001,
		Gender: "男",
	}
	data, _ := json.Marshal(stu)
	fmt.Println(string(data))

}
