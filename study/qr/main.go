package main

import (
	"image/color"
	"image/png"
	"net/url"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	// 原始JSON内容
	jsonStr := `{s(content, qrcode.Medium, 256)
	if err != nil {
		fmt.Printf("生成二维码失败: %v\n", err)
		return
	}

	// 将二维码数据写入文件
	err = os.WriteFile("artwork_qrcode.png", qrData, 0644)
	if err != nil {
		fmt.Printf("写入二维码文件失败: %v\n", err)
		return
	}}`

	// URL编码JSON参数
	encodedInfo := url.QueryEscape(jsonStr)

	// 构造完整URL
	content := "http:43.143.73.113/workinfo?info=" + encodedInfo

	// 生成二维码
	err := generateQRCode(content, "qrcode.png")
	if err != nil {
		panic(err)
	}
}

// generateQRCode 生成二维码并保存为PNG文件
func generateQRCode(content string, filename string) error {
	// 创建二维码
	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return err
	}

	// 自定义二维码样式
	qr.ForegroundColor = color.Black // 前景色
	qr.BackgroundColor = color.White // 背景色

	// 创建输出文件
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 编码为PNG并写入文件
	return png.Encode(file, qr.Image(256))
}
