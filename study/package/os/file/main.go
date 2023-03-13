package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file01, err := os.Create("./test01.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file01.Close()
	//for i := 0; i < 5; i++ {
	//	file.WriteString(fmt.Sprintln(i))
	//	file.Write([]byte("abc\n"))
	//}
	file, err := os.Open("./test.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	//var buf [128]byte
	//var content []byte
	//for {
	//	n, err := file01.Read(buf[:])
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	content = append(content, buf[:n]...)
	//}
	//fmt.Println(string(content))
	buf := make([]byte, 2046)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		file01.Write(buf[:n])
	}
	defer file01.Close()
}
