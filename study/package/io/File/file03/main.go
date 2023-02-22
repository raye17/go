package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	str := new(bytes.Buffer)
	str.WriteString("Hello,world!\n")
	str.WriteString("My name is sxy.\n")
	str.WriteString("What's your name?")
	file, err := os.Create("./test.md")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	if _, err := io.Copy(file, str); err != nil {
		fmt.Println("failed to copy:", err)
	}
	fmt.Println("file created")
}
