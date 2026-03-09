package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)

func main() {
	//	if err := joinPath(); err != nil {
	//		log.Fatal(err)
	//	}
	Path01()
}
func sTest05() {
	file, _ := os.Open("./ss.txt")
	buff := new(bytes.Buffer)
	_, err := io.Copy(buff, file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path.Ext(file.Name()))
	fmt.Println(buff.Bytes(), string(buff.Bytes()[0]))
	fmt.Println(buff.String())
}

// todo
func Path01() {
	fmt.Println(os.Getwd())
	s, err := os.Stat("./context/context.go")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Printf("s:%#v\n", s)
}
func joinPath() error {
	dir, _ := os.Getwd()
	p := dir + "/app/test"
	if err := os.MkdirAll(p, 0755); err != nil {
		return err
	}
	filePath := filepath.Join(p, "test.txt")
	return os.WriteFile(filePath, []byte("hello world"), 0644)
}
