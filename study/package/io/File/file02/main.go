package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	url := "../file01/proverbs.md"
	file, err := os.Open(url)
	if err != nil {
		fmt.Println("failed to open file: ", url)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			
		}
	}(file)
	p := make([]byte, 1024)
	for {
		n, err := file.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}
}
