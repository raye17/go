package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	content := "https://shimo.im/docs/zdkyBeeYPOUOZrA6"
	fileName := "QR.png" //输出文件名
	fileSize := 256      //二维码大小
	//生成二维码
	err := CreateIconPng(content, fileName, fileSize)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}
func CreateIconPng(content string, fileName string, fileSize int) (err error) {
	//生成一个二维码对象
	q, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return
	}
	//q.DisableBorder = true          //去除边框
	q.BackgroundColor = color.White //二维码背景颜色
	q.ForegroundColor = color.Black //二维码条纹颜色
	//设置二维码大小
	bgImg := q.Image(fileSize)

	//生成icon
	rgba, err := CreateIcon(bgImg)
	if err != nil {
		return
	}
	//合并二维码背景图和icon
	err = SaveImage(fileName, rgba)
	if err != nil {
		return
	}
	return
}
func CreateIcon(bgImg image.Image) (rgba *image.RGBA, err error) {
	//icon的路径
	iconPath := "img/icon.jpeg"
	avatarFile, err := os.Open(iconPath)
	if err != nil {
		return
	}
	//获得images
	avatarImg, err := jpeg.Decode(avatarFile)
	if err != nil {
		return
	}
	//设置icon大小
	avatarImg = resize.Resize(40, 40, avatarImg, resize.Lanczos3)

	//得到背景图的大小
	b := bgImg.Bounds()
	//居中设置icon到二维码图片
	offset := image.Pt((b.Max.X-avatarImg.Bounds().Max.X)/2, (b.Max.Y-avatarImg.Bounds().Max.Y)/2)
	rgba = image.NewRGBA(b)
	draw.Draw(rgba, b, bgImg, image.Point{X: 0, Y: 0}, draw.Src)
	draw.Draw(rgba, avatarImg.Bounds().Add(offset), avatarImg, image.Point{X: 0, Y: 0}, draw.Over)
	return
}
func SaveImage(p string, src image.Image) error {
	f, err := os.OpenFile(p, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	ext := filepath.Ext(p)
	if strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg") {
		err = jpeg.Encode(f, src, &jpeg.Options{Quality: 80})
	} else if strings.EqualFold(ext, ".png") {
		err = png.Encode(f, src)
	} else if strings.EqualFold(ext, ".gif") {
		err = gif.Encode(f, src, &gif.Options{NumColors: 256})
	}
	return err
}
