package main

import (
	"fmt"
	"io"
	"os"
)

type AlphaReader struct {
	reader io.Reader
}

func NewAlphaReader(src io.Reader) *AlphaReader {
	return &AlphaReader{reader: src}
}
func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}
func (a *AlphaReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}
	copy(p, buf)
	return n, nil
}
func main() {
	//reader := NewAlphaReader(strings.NewReader("hello,好，my name s raye,我的字，what's your name?"))
	f, err := os.Open("main.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	reader := NewAlphaReader(f)
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
}
