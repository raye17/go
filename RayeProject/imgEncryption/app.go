package main

import (
	"fmt"
	"os"
)

func main() {
	imagePath := "./data/1.png"
	srcZip := "./data/1.zip"
	outputPath := "./data/2.png"
	//outputPath := "D://2.png"
	// 打开图片文件
	imageFile, err := os.Open(imagePath)
	if err != nil {
		fmt.Printf("Cannot open the picture %s\n", imagePath)
		return
	}
	defer imageFile.Close()

	// 打开压缩包文件
	fFile, err := os.Open(srcZip)
	if err != nil {
		fmt.Printf("Cannot open the file %s\n", srcZip)
		return
	}
	defer fFile.Close()

	// 创建生成文件
	fFinish, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("Cannot create the file %s\n", outputPath)
		return
	}
	defer fFinish.Close()

	// 将图片内容复制到生成文件
	_, err = imageFile.WriteTo(fFinish)
	if err != nil {
		fmt.Printf("Error writing picture content: %v\n", err)
		return
	}

	// 将压缩包内容复制到生成文件
	// 改为使用专门的ZIP库处理
	zipData, err := os.ReadFile(srcZip)
	if err != nil {
		fmt.Printf("Error reading zip file: %v\n", err)
		return
	}

	// 确保写入完整的ZIP文件结构
	if _, err := fFinish.Write(zipData); err != nil {
		fmt.Printf("Error writing zip content: %v\n", err)
		return
	}

	fmt.Println("文件合成完成！")
}
