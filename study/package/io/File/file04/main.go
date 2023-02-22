package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("./test.md")
	if err != nil {
		fmt.Println(err, "failed to create file")
		os.Exit(1)
	}
	defer file.Close()
	if _, err := io.WriteString(file, "hello,my name is sxy!"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
