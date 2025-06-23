package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
)

var data = "H4sIAIr/QGgA/wVAsQrCMBDdA/mH9wlprAq3OQh2kUBQodtp0oJFD+6CgvjxZUg4lKLVjHAWDOnT44/IXU+7/SNQF7Y0RSL2LucTkmgjxOjdxaq++VUJKtK8S2z2FS2EebyHwJvjbf5Nma/L2JbnCtkYW7dnAAAA"

func main() {
	fmt.Println(len(data))
	s, err := gzips(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// gzips 先对输入的字符串进行base64解码，然后再用gzip解压缩
func gzips(encodedData string) (string, error) {
	// 1. base64解码
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return "", fmt.Errorf("base64解码失败: %w", err)
	}
	// 2. gzip解压缩
	reader, err := gzip.NewReader(bytes.NewReader(decodedData))
	if err != nil {
		return "", fmt.Errorf("创建gzip reader失败: %w", err)
	}
	defer reader.Close()

	// 读取解压后的数据
	decompressed, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("gzip解压失败: %w", err)
	}

	return string(decompressed), nil
}
