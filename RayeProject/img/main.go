package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"unicode/utf8"

	"github.com/golang/freetype"
	"github.com/skip2/go-qrcode"
	"github.com/xuri/excelize/v2"
	imageDraw "golang.org/x/image/draw"
	"golang.org/x/image/font"
)

func main() {
	Run()
}
func Run() {
	// 生成二维码
	content := "http://172.16.100.93:8006/workinfo?info=86d4d911-6474-477b-b0a3-3a9a1997b029"
	err := generateQRCode(content, "./data/qrcode.png")
	if err != nil {
		fmt.Printf("生成二维码失败: %v\n", err)
		return
	}
	basePath := "./data/展品标签.png"
	qrCodePath := "./data/qrcode.png"
	outputPath := fmt.Sprintf("./data/qrcode_styled.png")

	err = overlayQRCodeAndLabels(qrCodePath, basePath, outputPath)
	if err != nil {
		fmt.Printf("叠加二维码和标签失败: %v\n", err)
		return
	}

	// 导出到Excel
	excelPath := "./data/艺术品信息.xlsx"
	artworkName := "示例画作" // 替换为实际画作名称
	artistName := "示例画家"  // 替换为实际画家姓名
	if err := exportToExcel(outputPath, excelPath, artworkName, artistName); err != nil {
		fmt.Printf("导出到Excel失败: %v\n", err)
		return
	}
}
func overlayQRCodeAndLabels(qrCodePath, basePath, outputPath string) error {
	// 加载背景图片
	baseFile, err := os.Open(basePath)
	if err != nil {
		return fmt.Errorf("打开背景图片失败: %v", err)
	}
	defer baseFile.Close()

	baseImg, err := png.Decode(baseFile)
	if err != nil {
		return fmt.Errorf("解码背景图片失败: %v", err)
	}

	// 加载二维码图片
	qrFile, err := os.Open(qrCodePath)
	if err != nil {
		return fmt.Errorf("打开二维码图片失败: %v", err)
	}
	defer qrFile.Close()

	qrImg, err := png.Decode(qrFile)
	if err != nil {
		return fmt.Errorf("解码二维码图片失败: %v", err)
	}

	fmt.Println(":::")
	fmt.Println(baseImg.Bounds())
	fmt.Println(qrImg.Bounds())

	// 创建新图片
	bounds := baseImg.Bounds()
	baseWidth := bounds.Dx()
	baseHeight := bounds.Dy()
	fmt.Printf("基础图片尺寸: %dx%d\n", baseWidth, baseHeight)
	outImg := image.NewRGBA(bounds)

	// 绘制背景图片
	draw.Draw(outImg, bounds, baseImg, bounds.Min, draw.Over)

	// 叠加二维码
	// 二维码位置：左227px、上274px、右1127px、下273px
	// 二维码尺寸：宽度 = 1890 - 227 - 1127 = 536px，高度 = 1063 - 274 - 273 = 516px
	qrWidth := qrImg.Bounds().Dx()
	qrHeight := qrImg.Bounds().Dy()
	fmt.Println("qr:", qrWidth, qrHeight)

	// 定义目标区域(535x515)和居中位置
	targetWidth, targetHeight := 460, 480
	offsetX := 88 + (targetWidth-qrWidth)/2
	offsetY := 145 + (targetHeight-qrHeight)/2
	qrRect := image.Rect(offsetX, offsetY, offsetX+targetWidth, offsetY+targetHeight)
	qrCenterY := offsetY + targetHeight/2
	// 直接使用imageDraw进行缩放绘制
	imageDraw.CatmullRom.Scale(outImg, qrRect, qrImg, qrImg.Bounds(), draw.Over, nil)

	// 计算文字宽度
	c := freetype.NewContext()
	c.SetDPI(72)
	fontPath := "data/方正中雅宋_GBK.TTF" // 替换为实际字体路径
	fontBytes, _ := ioutil.ReadFile(fontPath)
	font, _ := freetype.ParseFont(fontBytes)
	c.SetFont(font)
	//c.SetFontSize(90) // 与后面设置的字体大小一致

	// 增加字符间距
	letterSpacing := 5.0 // 字间距，根据需要调整

	// 计算艺术家名字的宽度
	artistName := "于汪洋 李甜甜"
	artistFontSize := 70.0 // 根据高度调整字体大小
	artistWidth := 0
	for i, char := range artistName {
		if char == ' ' {
			// 空格宽度设为字体大小的1/4
			artistWidth += int(float64(artistFontSize) * 0.8)
		} else {
			glyphIndex := font.Index(char)
			hMetric := font.HMetric(c.PointToFixed(artistFontSize), glyphIndex)
			charWidth := float64(hMetric.AdvanceWidth) / 64.0
			artistWidth += int(charWidth)
			if i < utf8.RuneCountInString(artistName)-1 {
				artistWidth += int(letterSpacing)
			}
		}
	}
	fmt.Println("artistewitj ", artistWidth)
	artistX := 1300 - artistWidth/2
	artistY := qrCenterY - 110                        // 调整垂直位置
	artistColor := color.RGBA{0x00, 0x76, 0x86, 0xff} // #007686
	imgWithArtist, err := addLabels(outImg, artistName, artistX, artistY, artistColor, artistFontSize, fontPath, letterSpacing)
	if err != nil {
		return fmt.Errorf("添加作家名称失败: %v", err)
	}

	// 添加画作名称
	// 画作位置：上548px、左402px、右938px、下248px
	// 宽度 = 1890 - 402 - 938 = 550px，高度 = 1063 - 548 - 248 = 267px
	// 计算画作名字的宽度（加书名号）
	// 只计算第一行8个字符的宽度
	letterSpacing = 1.0
	artworkName := "家山树几颗"
	artworkName = fmt.Sprintf("《%s》", artworkName)
	artworkFontSize := 67.0
	totalWidth := 0
	for i, char := range artworkName {
		glyphIndex := font.Index(char)
		hMetric := font.HMetric(c.PointToFixed(artworkFontSize), glyphIndex)
		charWidth := float64(hMetric.AdvanceWidth) / 64.0
		totalWidth += int(charWidth)
		if i < utf8.RuneCountInString(artworkName)-1 {
			totalWidth += int(letterSpacing)
		}
	}
	artworkX := 1300 - totalWidth/2
	artworkY := qrCenterY + 10
	artworkColor := color.RGBA{0x00, 0x76, 0x86, 0xff} // #007686
	finalImg, err := addLabels(imgWithArtist, artworkName, artworkX, artworkY, artworkColor, artworkFontSize, fontPath, letterSpacing)
	if err != nil {
		return fmt.Errorf("添加画作名称失败: %v", err)
	}

	// 保存最终图片
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("创建输出文件失败: %v", err)
	}
	// 压缩图片质量 (0-100, 数值越小压缩率越高)
	quality := 25
	opt := jpeg.Options{Quality: quality}
	if err := jpeg.Encode(outFile, finalImg, &opt); err != nil {
		return fmt.Errorf("压缩图片失败: %v", err)
	}

	defer outFile.Close()
	return nil

	// if err := png.Encode(outFile, finalImg); err != nil {
	// 	return fmt.Errorf("编码图片失败: %v", err)
	// }

}
func generateQRCode(content, filePath string) error {
	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return err
	}

	qr.DisableBorder = true
	img := qr.Image(150)

	// 创建透明背景的新图像
	transparentImg := image.NewRGBA(img.Bounds())
	draw.Draw(transparentImg, transparentImg.Bounds(), image.Transparent, image.Point{}, draw.Src)

	// 只复制黑色像素
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			// 如果是黑色像素则保留
			if r == 0 && g == 0 && b == 0 {
				transparentImg.Set(x, y, color.Black)
			}
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, transparentImg)
}
func addLabels(img image.Image, label string, x, y int, fontColor color.Color, size float64, fontPath string, letterSpacing float64) (image.Image, error) {
	bound := img.Bounds()
	// 创建一个新的图片
	rgba := image.NewRGBA(image.Rect(0, 0, bound.Dx(), bound.Dy()))
	// 读取字体
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return rgba, err
	}
	myFont, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return rgba, err
	}

	draw.Draw(rgba, rgba.Bounds(), img, bound.Min, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(myFont)
	c.SetFontSize(size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	uni := image.NewUniform(fontColor)
	c.SetSrc(uni)
	c.SetHinting(font.HintingNone)
	// 分行处理
	//lineHeight := int(c.PointToFixed(size*1.0) >> 6) // 行高为字体大小的1.5倍
	currentY := y
	lineChars := 0
	// 手动绘制每个字符以控制间距
	// 逐字符绘制，控制字间距
	currentX := float64(x)
	for _, char := range label {
		// 将字符转换为字符串
		charStr := string(char)

		// // 换行处理
		// if lineChars >= 8 && charStr != "》" { // 书名号不换行
		// 	currentX = float64(x)
		// 	currentY += lineHeight
		// 	lineChars = 0
		// }

		pt := freetype.Pt(int(currentX), currentY+int(c.PointToFixed(size)>>6))
		_, err := c.DrawString(charStr, pt)
		if err != nil {
			return rgba, fmt.Errorf("绘制字符失败: %v", err)
		}

		glyphIndex := myFont.Index(char)
		hMetric := myFont.HMetric(c.PointToFixed(size), glyphIndex)
		charWidth := float64(hMetric.AdvanceWidth) / 64.0

		currentX += charWidth + letterSpacing
		lineChars++
	}

	return rgba, nil
}
func exportToExcel(imagePath, excelPath, artworkName, artistName string) error {
	f := excelize.NewFile()
	defer f.Close()

	// 设置表头
	// 设置表头样式
	style, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	f.SetCellStyle("Sheet1", "A1", "C1", style)
	f.SetCellValue("Sheet1", "A1", "画作名称")
	f.SetCellValue("Sheet1", "B1", "画家姓名")
	f.SetCellValue("Sheet1", "C1", "带二维码的图片")

	// 设置行高和列宽
	f.SetRowHeight("Sheet1", 1, 30)         // 表头行高
	f.SetRowHeight("Sheet1", 2, 80)         // 数据行高
	f.SetColWidth("Sheet1", "A", "B", 20)   // 文字列宽
	f.SetColWidth("Sheet1", "C", "C", 22.5) // 图片列宽

	// 添加数据
	f.SetCellValue("Sheet1", "A2", artworkName)
	f.SetCellValue("Sheet1", "B2", artistName)
	f.SetCellStyle("Sheet1", "A2", "B2", style)

	// 插入图片并调整大小
	if err := f.AddPicture("Sheet1", "C2", imagePath, &excelize.GraphicOptions{
		ScaleX:  0.08, // 水平缩放比例
		ScaleY:  0.08, // 垂直缩放比例
		OffsetX: 5,    // 水平偏移
		OffsetY: 5,    // 垂直偏移
	}); err != nil {
		return fmt.Errorf("插入图片失败: %v", err)
	}

	// 保存Excel文件
	if err := f.SaveAs(excelPath); err != nil {
		return fmt.Errorf("保存Excel失败: %v", err)
	}

	return nil
}
