package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"study/model"
	"study/service"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func main() {
	r := gin.Default()
	r.POST("/account", ReadArtistAccountInfo)
	r.Run(":8081")
}
func ReadArtistAccountInfo(c *gin.Context) {
	// 1. 上传文件
	excelFile, err := c.FormFile("excel")
	if err != nil {
		c.JSON(400, gin.H{"error": "缺少 Excel 文件 excel"})
		return
	}
	// 2. 保存临时文件
	tempDir := "tmp"
	os.MkdirAll(tempDir, 0755)
	excelPath := filepath.Join(tempDir, "artists.xlsx")
	if err = c.SaveUploadedFile(excelFile, excelPath); err != nil {
		c.JSON(500, gin.H{"error": "保存 Excel 失败"})
		return
	}
	defer os.RemoveAll(tempDir)
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		return
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return
	}
	log.Println("start read excel...")
	var artists []model.ArtistAccount
	for i, row := range rows {
		if i == 0 || len(row) < 2 {
			continue
		}
		tmp := model.ArtistAccount{
			Account: make(map[uint64]string),
			Name:    strings.TrimSpace(row[1]),
			SubNum:  strings.TrimSpace(row[2]),
		}
		youtube, _ := f.GetCellValue(sheetName, fmt.Sprintf("C%d", i+1))
		if youtube != "" {
			tmp.Account[2] = strings.TrimSpace(youtube)

		}
		ins, _ := f.GetCellValue(sheetName, fmt.Sprintf("D%d", i+1))
		if ins != "" {
			tmp.Account[3] = strings.TrimSpace(ins)

		}
		tiktok, _ := f.GetCellValue(sheetName, fmt.Sprintf("E%d", i+1))
		if tiktok != "" {
			tmp.Account[1] = strings.TrimSpace(tiktok)

		}
		artists = append(artists, tmp)
	}
	service.Success(c, "success", artists)
}
