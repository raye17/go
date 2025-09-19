package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	oldName := "1.mp4"
	newName := "tmp/oss/9.mp4"

	// 创建目标目录
	dirPath := filepath.Dir(newName)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		log.Println("创建目录失败:", err)
		return
	}

	log.Println("创建目录成功:", dirPath)

	// 检查源文件是否存在
	if _, err := os.Stat(oldName); os.IsNotExist(err) {
		log.Println("源文件不存在:", oldName)
		return
	}

	// 打开新文件用于复制
	srcFile, err := os.Open(newName)
	if err != nil {
		log.Println("打开文件失败:", err)
		return
	}
	defer srcFile.Close()

	// 构造上级目录路径
	dstPath := filepath.Join("..", newName)
	// 创建上级目录
	dstDir := filepath.Dir(dstPath)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		log.Println("创建上级目录失败:", err)
		return
	}
	dstFile, err := os.Create(dstPath)
	if err != nil {
		log.Println("创建目标文件失败:", err)
		return
	}
	defer dstFile.Close()

	// 执行复制
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		log.Println("复制失败:", err)
		return
	}

	log.Println("复制成功:", dstPath)
}
