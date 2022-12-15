package main

import (
	"fmt"
	"net/http"
)

var url = []string{
	"https://www.baidu.com",
	"https://www.google.com",
	"https://www.taobao.com",
}

func main() {
	for k, v := range url {
		fmt.Println(k)
		res, err := http.Head(v)
		if err != nil {
			fmt.Printf("请求%s失败, 失败原因：%v\n", v, err)
			continue
		}
		fmt.Printf("请求%s成功，状态为：%v\n", v, res.Status)
	}
}
