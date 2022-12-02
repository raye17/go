package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!,err:", err)
		return
	} else {
		fmt.Println("open file success!")

	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	var temp = make([]byte, 480)
	n, err := file.Read(temp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("read success")
	}
	fmt.Println(n)
	fmt.Println(string(temp))
}
