package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/xuri/excelize/v2"
)

//	func main() {
//		redis := cache.NewClientRedis()
//		fmt.Println(redis)
//		keys, err := redis.Keys(context.TODO(), "*").Result()
//		if err != nil {
//			fmt.Println(err)
//		}
//		for _, key := range keys {
//			typeStr, err := redis.Type(context.TODO(), key).Result()
//			if err != nil {
//				fmt.Println("获取类型错误:", err)
//				continue
//			}
//			switch typeStr {
//			case "string":
//				v, err := redis.Get(context.TODO(), key).Result()
//				if err != nil {
//					fmt.Println("string类型获取失败:", err)
//				} else {
//					fmt.Println("[string] k:", key, "v:", v)
//				}
//			case "zset":
//				zvals, err := redis.ZRangeWithScores(context.TODO(), key, 0, -1).Result()
//				if err != nil {
//					fmt.Println("zset类型获取失败:", err)
//				} else {
//					fmt.Printf("[zset] k: %s v: %v\n", key, zvals)
//				}
//			case "hash":
//				hvals, err := redis.HGetAll(context.TODO(), key).Result()
//				if err != nil {
//					fmt.Println("hash类型获取失败:", err)
//				} else {
//					fmt.Printf("[hash] k: %s v: %v\n", key, hvals)
//				}
//			default:
//				fmt.Println("未处理类型:", typeStr, "key:", key)
//			}
//		}
//	}
func main() {
	fileDir := "./runtime/import/"
	filename := "画家视频详情记录.xlsx"
	filePath := filepath.Join(fileDir, filename)
	_ = os.MkdirAll(fileDir, os.ModePerm)

	var f *excelize.File
	sheet := "Sheet1"

	// 判断文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 文件不存在，新建文件和Sheet
		f = excelize.NewFile()
		f.SetSheetName(f.GetSheetName(0), sheet)

		// 写表头
		headers := []string{"序号", "画家名", "标题", "uuid", "youtube", "instagram", "tiktok"}
		for col, h := range headers {
			_ = f.SetCellValue(sheet, string('A'+col)+"1", h)
		}
	} else {
		// 文件存在，打开
		var err error
		f, err = excelize.OpenFile(filePath)
		if err != nil {
			return
		}
	}

	// 找到最后一行，追加数据
	rows, err := f.GetRows(sheet)
	if err != nil {
		return
	}

	// 计算下一行，从表头之后开始
	startRow := len(rows) + 1
	if startRow == 1 {
		startRow = 2 // 文件新建或没有数据，从第2行开始
	}
	var artistInfos = []ArtistVideoDetail{
		{
			ArtistName: "画家3",
			Title:      "作品3",
			WorkUuid:   "uuid3",
			Youtube:    "https://www.youtube.com",
			Instagram:  "https://www.instagram.com",
			TikTok:     "https://www.tiktok.com",
		},
		{
			ArtistName: "画家4",
			Title:      "作品4",
			WorkUuid:   "uuid4",
			Youtube:    "https://www.youtube.com",
			Instagram:  "https://www.instagram.com",
			TikTok:     "https://www.tiktok.com",
		},
	}

	// 写数据
	for i, artistInfo := range artistInfos {
		row := startRow + i
		_ = f.SetCellValue(sheet, "A"+strconv.Itoa(row), row-1) // 序号连续
		_ = f.SetCellValue(sheet, "B"+strconv.Itoa(row), artistInfo.ArtistName)
		_ = f.SetCellValue(sheet, "C"+strconv.Itoa(row), artistInfo.Title)
		_ = f.SetCellValue(sheet, "D"+strconv.Itoa(row), artistInfo.WorkUuid)
		_ = f.SetCellValue(sheet, "E"+strconv.Itoa(row), artistInfo.Youtube)
		_ = f.SetCellValue(sheet, "F"+strconv.Itoa(row), artistInfo.Instagram)
		_ = f.SetCellValue(sheet, "G"+strconv.Itoa(row), artistInfo.TikTok)
	}
	if err := f.SaveAs(filePath); err != nil {
		fmt.Println("save excel err: ", err)
		return
	}

}

type ArtistVideoDetail struct {
	Id         string `json:"id"`
	ArtistName string `json:"artistName"`
	Title      string `json:"title"`
	WorkUuid   string `json:"workUuid"`
	Youtube    string `json:"youtube"`
	Instagram  string `json:"instagram"`
	TikTok     string `json:"tiktok"`
}
