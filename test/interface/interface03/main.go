package main

import "fmt"

func main() {
	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "raye"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Printf("%v", studentInfo)
	//gin框架的gin.H{}
	fmt.Println()
	s := fmt.Sprintf("字符串%s", "string")
	fmt.Println(s)
}
