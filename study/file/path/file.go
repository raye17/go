package path

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func sTest05() {
	file, _ := os.Open("./ss.txt")
	buff := new(bytes.Buffer)
	_, err := io.Copy(buff, file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path.Ext(file.Name()))
	fmt.Println(buff.Bytes(), string(buff.Bytes()[0]))
	fmt.Println(string(buff.Bytes()))
}

// todo
func Path01() {
	fmt.Println(os.Getwd())
	s, err := os.Stat("./context/context.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", s)
}
