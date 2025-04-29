package main

import (
	"fmt"
	"os"
)

func main() {
	imagePath := "./data/1.png"
	fileToHide := "./data/1.zip"
	outputPath := "./data/2.png"

	// 打开图片文件
	fPic, err := os.Open(imagePath)
	if err != nil {
		fmt.Printf("Cannot open the picture %s\n", imagePath)
		return
	}
	defer fPic.Close()

	// 打开压缩包文件
	fFile, err := os.Open(fileToHide)
	if err != nil {
		fmt.Printf("Cannot open the file %s\n", fileToHide)
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
	_, err = fPic.WriteTo(fFinish)
	if err != nil {
		fmt.Printf("Error writing picture content: %v\n", err)
		return
	}

	// 将压缩包内容复制到生成文件
	_, err = fFile.WriteTo(fFinish)
	if err != nil {
		fmt.Printf("Error writing file content: %v\n", err)
		return
	}

	fmt.Println("文件合成完成！")
}
