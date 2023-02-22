package main

import (
	"fmt"
	"io"
)

type AlphaReader struct {
	src string //资源
	cur int    //当前读取到的位置
}

func NewAlphaReader(src string) *AlphaReader {
	return &AlphaReader{src: src}
}
func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}
func (a *AlphaReader) Read(p []byte) (int, error) {
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}
	x := len(a.src) - a.cur
	n, bound := 0, 0
	if x >= len(p) {
		bound = len(p)
	} else if x < len(p) {
		bound = x
	}
	buf := make([]byte, bound)
	for n < bound {
		if char := alpha(a.src[a.cur]); char != 0 {
			buf[n] = char
		}
		n++
		a.cur++
	}
	copy(p, buf)
	return n, nil
}
func main() {
	reader := NewAlphaReader("hello,好，my name s raye,我的字，what's your name?")
	fmt.Println(len(reader.src))
	p := make([]byte, 23)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}
}
