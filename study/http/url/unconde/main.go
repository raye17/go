package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"image/png"
	"net/url"
	"os"

	"github.com/skip2/go-qrcode"
)

type ArtworkInfo struct {
	Title        string `json:"title"`
	Image        string `json:"image"`
	Artist       string `json:"artist"`
	Size         string `json:"size"`
	Year         string `json:"year"`
	Type         string `json:"type"`
	Material     string `json:"material"`
	Appreciation string `json:"appreciation"`
}

const (
	QrCodeDevBaseURL = "http://172.16.100.93:8004"
)

func main() {
	info := ArtworkInfo{
		Title:        "家山树几棵",
		Image:        "https://e-cdn.fontree.cn/exhibition/prod/default/国展报名/0737054f-b2c7-4c92-8e75-1100faca19da.png",
		Artist:       "李甜甜02",
		Size:         "2400cmX54656cm",
		Year:         "",
		Type:         "中国画",
		Material:     "宣纸",
		Appreciation: "",
	}

	jsonData, err := json.Marshal(info)
	if err != nil {
		fmt.Printf("序列化失败: %v\n", err)
		return
	}
	//baseURL := QrCodeDevBaseURL
	encodedInfo := url.QueryEscape(string(jsonData))

	// 生成二维码内容
	content := fmt.Sprintf("%s/workinfo?info=%s", QrCodeDevBaseURL, encodedInfo)

	// 生成二维码图片
	qrData, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		fmt.Printf("生成二维码失败: %v\n", err)
		return
	}

	// 将二维码数据写入文件
	err = os.WriteFile("artwork_qrcode.png", qrData, 0644)
	if err != nil {
		fmt.Printf("写入二维码文件失败: %v\n", err)
		return
	}

	fmt.Println("二维码已生成: artwork_qrcode.png")
	//fmt.Println("访问链接:", content)
}

// generateQRCode 生成二维码并保存为PNG文件
func generateQRCode(content string, filename string) error {
	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return err
	}

	qr.ForegroundColor = color.Black
	qr.BackgroundColor = color.White

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, qr.Image(256))
}
