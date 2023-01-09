package main

import "fmt"

func main() {
	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "raye"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	s1 := studentInfo["name"]
	v, ok := s1.(int)
	fmt.Println(v, ok)
	fmt.Printf("%v\n", studentInfo)
	s := fmt.Sprintf("字符串%s", "string")
	fmt.Println(s)
}
