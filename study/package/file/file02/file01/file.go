package main

import (
	"fmt"
	"io"
	"os"
)

func readAll() {
	fileObj, err := os.Open("../test01.txt")
	if err != nil {
		fmt.Printf("open file failed,err%v\n", err)
		return
	} else {
		fmt.Println("open success!")
	}
	defer fileObj.Close()
	for {
		var temp = make([]byte, 18888)
		n, err := fileObj.Read(temp)
		if err == io.EOF {
			fmt.Println(string(temp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v\n", err)
			return
		}
		fmt.Printf("read %d bytes from file.\n", n)
		fmt.Println(string(temp[:n]))
	}
}

func main() {
	readAll()
}
