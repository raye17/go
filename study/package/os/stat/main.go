package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)
	dir = filepath.Join(dir, "main.go")
	i, err := os.Stat(dir)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(i.IsDir())
}
