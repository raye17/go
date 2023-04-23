package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

func main() {
	s := "my name is sxy,and i like novel"
	k := bytes.NewBuffer([]byte(s))
	fmt.Println([]byte(s))
	fmt.Println(k)
	fmt.Printf("k:%v\n", k)
	gz, err := gzip.NewReader(k)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(gz)
}
