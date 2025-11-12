package fiee

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

type WorkWithSubNum struct {
	ArtistName  string
	UpdateTime  string
	ArtistPhone string
	WorkUUID    string
	Title       string
	SubNum      string
}

// 假设 db1 是 cast_work_log + cast_work
// db2 是 user 表
var results []WorkWithSubNum

// 1️⃣ 查询 cast_work_log + cast_work
type WorkLogTmp struct {
	WorkUUID    string
	ArtistUUID  string
	ArtistName  string
	UpdateTime  string
	Title       string
	ArtistPhone string
}

func copyConn(a, b net.Conn) {
	defer a.Close()
	defer b.Close()
	_, _ = io.Copy(a, b)
}
func readArtistVideoInfo(excelPath string) (map[string]ArtistMedia, error) {
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}
	log.Println("start read excel")
	var artists map[string]ArtistMedia
	for i, row := range rows {
		if i == 0 || len(row) < 2 {
			continue
		}
		if i == 165 {
			break
		}
		id, _ := f.GetCellValue(sheetName, fmt.Sprintf("A%d", i+1))
		if id != "" {
			id = strings.TrimSpace(id)
		}
		artistName, _ := f.GetCellValue(sheetName, fmt.Sprintf("B%d", i+1))
		if artistName != "" {
			artistName = strings.TrimSpace(artistName)
		}
		title, _ := f.GetCellValue(sheetName, fmt.Sprintf("C%d", i+1))
		if title != "" {
			title = strings.TrimSpace(title)
		}
		uuid, _ := f.GetCellValue(sheetName, fmt.Sprintf("D%d", i+1))
		if uuid != "" {
			uuid = strings.TrimSpace(uuid)
		}
		youtube, _ := f.GetCellValue(sheetName, fmt.Sprintf("E%d", i+1))
		if youtube != "" {
			youtube = strings.TrimSpace(youtube)
		}
		instagram, _ := f.GetCellValue(sheetName, fmt.Sprintf("F%d", i+1))
		if instagram != "" {
			instagram = strings.TrimSpace(instagram)
		}
		tiktok, _ := f.GetCellValue(sheetName, fmt.Sprintf("G%d", i+1))
		if tiktok != "" {
			tiktok = strings.TrimSpace(tiktok)
		}
		if youtube == "账号待解封" {
			youtube = "2006-01-02 15:04:05"
		}
		if instagram == "账号待解封" {
			instagram = "2006-01-02 15:04:05"
		}
		if tiktok == "账号待解封" {
			tiktok = "2006-01-02 15:04:05"
		}
		youtubeTime := parseTime(youtube)
		instagramTime := parseTime(instagram)
		tiktokTime := parseTime(tiktok)
		maxTime := maxTime(youtubeTime, instagramTime, tiktokTime)
		artist := ArtistMedia{
			Id:         id,
			Name:       artistName,
			Title:      title,
			Uuid:       uuid,
			Youtube:    youtube,
			Instagram:  instagram,
			TikTok:     tiktok,
			UpdateTime: maxTime.Format("2006-01-02 15:04:05"),
		}
		if artists == nil {
			artists = make(map[string]ArtistMedia)
		}
		artists[id] = artist
	}
	return artists, nil
}

type ArtistMedia struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	Img        string `json:"img"`
	Video      string `json:"video"`
	Uuid       string `json:"uuid"`
	Youtube    string `json:"youtube"`
	Instagram  string `json:"instagram"`
	TikTok     string `json:"tiktok"`
	UpdateTime string `json:"update_time"`
}

func parseTime(str string) *time.Time {
	if str == "" {
		return nil
	}

	// 支持多种格式尝试解析
	formats := []string{
		"2006-01-02 15:04:05",
		"2006-01-02",
		"2006/01/02 15:04:05",
		"2006/1/2",
		"2006/01/02",
		"2006-1-2",
		"2006/1/2 15:04:05",
	}

	for _, layout := range formats {
		if t, err := time.ParseInLocation(layout, str, time.Local); err == nil {
			return &t
		}
	}

	log.Println("parse time err: unsupported format:", str)
	return nil
}
func readArtistDetail(excelPath string) (map[string]string, error) {
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}
	log.Println("start read excel")
	var artistsVideo map[string]string
	for i, row := range rows {
		if i == 0 || i == 1 || len(row) < 2 {
			continue
		}
		if i == 165 {
			break
		}
		id, _ := f.GetCellValue(sheetName, fmt.Sprintf("A%d", i+1))
		if id != "" {
			id = strings.TrimSpace(id)
		}

		uuid, _ := f.GetCellValue(sheetName, fmt.Sprintf("D%d", i+1))
		if uuid != "" {
			uuid = strings.TrimSpace(uuid)
		}
		if artistsVideo == nil {
			artistsVideo = make(map[string]string)
		}
		artistsVideo[id] = uuid
	}
	return artistsVideo, nil
}
func maxTime(times ...*time.Time) *time.Time {
	var max *time.Time
	for _, t := range times {
		if t == nil {
			continue
		}
		if max == nil || t.After(*max) {
			max = t
		}
	}
	return max
}
func exportPublishRecordsToExcel(artistInfos []WorkWithSubNum) (string, error) {
	fileDir := "./import/"
	filename := "画家视频时间0925.xlsx"
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
		headers := []string{"视频主题", "确认时间", "画家名字", "手机号", "用户编号"}
		for col, h := range headers {
			_ = f.SetCellValue(sheet, string('A'+col)+"1", h)
		}
	} else {
		// 文件存在，打开
		var err error
		f, err = excelize.OpenFile(filePath)
		if err != nil {
			return "", err
		}
	}

	// 找到最后一行，追加数据
	rows, err := f.GetRows(sheet)
	if err != nil {
		return "", err
	}

	// 计算下一行，从表头之后开始
	startRow := len(rows) + 1
	if startRow == 1 {
		startRow = 2 // 文件新建或没有数据，从第2行开始
	}

	// 写数据
	for i, artistInfo := range artistInfos {
		row := startRow + i
		var update_time string
		if !strings.HasSuffix(artistInfo.UpdateTime, "00:00:00") {
			update_time = strings.Replace(artistInfo.UpdateTime, "00:00:00", "", -1)
		} else {
			update_time = artistInfo.UpdateTime
		}

		//_ = f.SetCellValue(sheet, "A"+strconv.Itoa(row), row-1) // 序号连续
		_ = f.SetCellValue(sheet, "A"+strconv.Itoa(row), artistInfo.Title)
		_ = f.SetCellValue(sheet, "B"+strconv.Itoa(row), update_time)
		_ = f.SetCellValue(sheet, "C"+strconv.Itoa(row), artistInfo.ArtistName)
		_ = f.SetCellValue(sheet, "D"+strconv.Itoa(row), artistInfo.ArtistPhone)
		_ = f.SetCellValue(sheet, "E"+strconv.Itoa(row), artistInfo.SubNum)
	}

	// 保存文件
	if err = f.SaveAs(filePath); err != nil {
		fmt.Println("saveAs err: ", err)
		return "", err
	}

	return filePath, nil
}
