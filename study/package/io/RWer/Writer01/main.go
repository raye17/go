package main

import (
	"bytes"
	"fmt"
	"os"
)

// 使用bytes.Buffer类型作为io.Writer将数据写入内存缓冲区
func main() {
	proverbs := []string{
		"hello,my name is sxy",
		"lcx is a pig",
		"你好，世界",
		"嘿嘿，哈哈，heihei",
	}
	var writer bytes.Buffer
	for _, p := range proverbs {
		p = p + "\n"
		n, err := writer.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
	fmt.Println(writer.String())
	fmt.Println(writer)
}
