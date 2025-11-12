package common

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/xuri/excelize/v2"
)

func renameFile() {
	dir := "E://待上传 - 副本//fiee3 - 副本//9.26"
	excelPath := dir + "/9.26.xlsx"
	dataDir := dir

	// 打开 Excel
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 假设数据在第一个sheet
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		panic(err)
	}

	// 遍历每一行，跳过表头
	for i, row := range rows {
		if i == 0 || i == 1 {
			continue
		}
		if len(row) < 3 {
			continue
		}

		seq := strings.TrimSpace(row[0])
		artist := strings.TrimSpace(row[1])
		title := strings.TrimSpace(row[2])
		subNum, err := f.GetCellValue(f.GetSheetName(0), fmt.Sprintf("D%d", i+1))
		if err != nil {
			log.Println(err)
			return
		}
		subNum = strings.TrimSpace(subNum)
		artistDir := artist
		if subNum != "" {
			artistDir = artist + subNum
		}
		artistDir = filepath.Join(dataDir, artist)
		if _, err := os.Stat(artistDir); os.IsNotExist(err) {
			fmt.Printf("❌ 没找到画家目录: %s\n", artistDir)
			continue
		}

		// jpg 和 mp4 两种
		// 格式分组：图片组 / 视频组
		groups := [][]string{
			{".jpg", ".png"}, // 图片类
			{".mp4", ".mov"}, // 视频类
		}

		for _, exts := range groups {
			found := false
			for _, ext := range exts {
				oldPath := filepath.Join(artistDir, title+ext)
				newPath := filepath.Join(artistDir, seq+ext)

				if _, err := os.Stat(oldPath); err == nil {
					if err := os.Rename(oldPath, newPath); err != nil {
						fmt.Printf("❌ 重命名失败 %s -> %s: %v\n", oldPath, newPath, err)
					} else {
						//fmt.Printf("✅ %s -> %s\n", oldPath, newPath)
					}
					found = true
					break // 找到一个就跳出当前组，避免重复
				}
			}

			if !found {
				fmt.Printf("⚠️ 没找到文件: %s %s (支持扩展: %v)\n", artist, title, exts)
			}
		}
	}
}
func exportFilesToExcel(dataDir, excelPath string) {
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetSheetName(f.GetSheetName(0), sheet)

	// 写表头
	f.SetCellValue(sheet, "A1", "序号")
	f.SetCellValue(sheet, "B1", "画家")
	f.SetCellValue(sheet, "C1", "标题")
	f.SetCellValue(sheet, "D1", "用户编号")

	rowNum := 2
	seq := 1

	// 遍历画家目录（按目录顺序，不排序）
	artistDirs, _ := os.ReadDir(dataDir)
	for _, ad := range artistDirs {
		if !ad.IsDir() {
			continue
		}

		artistName := ad.Name()
		userCode := ""
		if idx := strings.Index(artistName, "FE"); idx > 0 {
			userCode = artistName[idx:]
			artistName = artistName[:idx]
		}

		artistPath := filepath.Join(dataDir, ad.Name())

		// 收集文件
		files, _ := os.ReadDir(artistPath)
		mp4Files := []string{}
		jpgFiles := []string{}
		for _, f := range files {
			ext := strings.ToLower(filepath.Ext(f.Name()))
			if ext == ".mp4" || ext == ".mov" {
				mp4Files = append(mp4Files, f.Name())
			} else if ext == ".jpg" || ext == ".png" {
				jpgFiles = append(jpgFiles, f.Name())
			}
		}

		// 不再排序，按文件系统顺序
		n := len(mp4Files)
		if len(jpgFiles) < n {
			n = len(jpgFiles)
		}

		for i := 0; i < n; i++ {
			title := strings.TrimSuffix(jpgFiles[i], filepath.Ext(jpgFiles[i])) // 从 JPG 提取标题

			f.SetCellValue(sheet, fmt.Sprintf("A%d", rowNum), seq)
			f.SetCellValue(sheet, fmt.Sprintf("B%d", rowNum), artistName)
			f.SetCellValue(sheet, fmt.Sprintf("C%d", rowNum), title)
			f.SetCellValue(sheet, fmt.Sprintf("D%d", rowNum), userCode)

			rowNum++
			seq++
		}
	}

	if err := f.SaveAs(excelPath); err != nil {
		log.Fatal("保存 Excel 失败:", err)
	}

	fmt.Println("✅ Excel 导出完成，路径:", excelPath)
}
