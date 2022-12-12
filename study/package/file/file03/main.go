package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("readme.md")
	if err != nil {
		fmt.Println(err)
		file, err = os.Create("readme.md")
		if err != nil {
			fmt.Errorf("error:%v", err)
		}
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println(line)
			return
		}
		if err != nil {
			fmt.Println("error:", err)
		}
		//fmt.Println(prefix)
		fmt.Print(line)
	}

}
