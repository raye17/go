package main

import (
	"context"
	"encoding/json"
	"fiee/import/work/common"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type UploadResult struct {
	FilePath string
	Success  bool
	Err      error
}

type ApiResponse struct {
	Status int         `json:"status"` // 0 成功，1 失败
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Code   int         `json:"code"`
	Error  interface{} `json:"error"`
	Err    string      `json:"err"`
}

func main() {
	apiURL := "https://erpapi.fiee.com/api/fiee/import/data/publish"
	token := ""
	dir := "E://待上传 - 副本//28新加//"
	excelPath := dir + "28遗漏.xlsx"
	baseDir := dir + "28排查遗漏 - 副本"
	resultsFile := "upload_results.json"
	results := []UploadResult{}
	// 读取子目录
	subDirs, err := os.ReadDir(baseDir)
	if err != nil {
		log.Fatal(err)
	}
	var dirNames []string
	fmt.Println(subDirs)
	fmt.Println("***************")
	for _, d := range subDirs {
		if d.IsDir() {
			dirNames = append(dirNames, d.Name())
		}
	}
	for k, v := range dirNames {
		fmt.Println(k, v)
	}
	sort.Strings(dirNames)
	for k, v := range dirNames {
		fmt.Println(k, v)
	}
	// 已经上传过的
	uploaded := loadUploaded(resultsFile)

	// channel 传输压缩结果
	type ZipTask struct {
		Name string
		Path string
		Err  error
	}
	zipCh := make(chan ZipTask, len(dirNames))

	var wg sync.WaitGroup

	// 并发压缩

	for idx, name := range dirNames {
		wg.Add(1)
		go func(name string, index int) {
			defer wg.Done()

			origFolderPath := filepath.Join(baseDir, name) // 原始画家文件夹
			tempDir := filepath.Join(baseDir, "tmp_"+strconv.Itoa(index))
			topDir := filepath.Join(tempDir, strconv.Itoa(index))
			destDir := filepath.Join(topDir, name) // 临时目录里的画家文件夹

			_ = os.RemoveAll(tempDir)
			_ = os.MkdirAll(destDir, os.ModePerm)

			if err := common.CopyDir(origFolderPath, destDir); err != nil {
				zipCh <- ZipTask{Name: name, Err: err}
				return
			}

			zipPath := filepath.Join(baseDir, strconv.Itoa(index)+".zip")

			if err := common.ZipFolder(topDir, zipPath); err != nil {
				zipCh <- ZipTask{Name: name, Err: err}
				return
			}

			zipCh <- ZipTask{Name: name, Path: zipPath}
		}(name, idx+1) // idx+1 作为编号
	}

	// 另一个 goroutine 等待所有压缩完成后关闭 channel
	go func() {
		wg.Wait()
		close(zipCh)
	}()
	// 打开结果文件（追加写）
	f, _ := os.OpenFile(resultsFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer f.Close()

	// 上传 goroutine 循环 channel
	for task := range zipCh {
		zipPath := task.Path
		if zipPath == "" {
			zipPath = filepath.Join(baseDir, task.Name+".zip")
		}

		if uploaded[zipPath] {
			log.Println("⏩ 已上传过，跳过:", zipPath)
			continue
		}

		if task.Err != nil {
			f.WriteString(fmt.Sprintf("%s -> 压缩失败: %v\n", zipPath, task.Err))
			results = append(results, UploadResult{FilePath: zipPath, Success: false, Err: task.Err})
			continue
		}

		log.Println("🚀 开始上传:", zipPath)
		upErr := uploadStream(apiURL, token, excelPath, zipPath)

		if upErr != nil {
			log.Println("❌ 上传失败:", zipPath, upErr)
			f.WriteString(fmt.Sprintf("%s -> 上传失败: %v\n", zipPath, upErr))
			results = append(results, UploadResult{FilePath: zipPath, Success: false, Err: upErr})
		} else {
			log.Println("✅ 上传成功:", zipPath)
			f.WriteString(fmt.Sprintf("%s -> 成功\n", zipPath))
			results = append(results, UploadResult{FilePath: zipPath, Success: true})
		}

		//_ = os.Remove(zipPath)
	}

	fmt.Println("🎉 全部任务完成")
}

// 流式上传，避免一次性占用内存
func uploadStream(apiURL, token, excelPath, zipPath string) error {
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	// 在单独 goroutine 写入数据
	go func() {
		defer pw.Close()
		defer writer.Close()

		// excel 文件
		excelFile, err := os.Open(excelPath)
		if err != nil {
			pw.CloseWithError(err)
			return
		}
		defer excelFile.Close()

		excelPart, err := writer.CreateFormFile("excel", filepath.Base(excelPath))
		if err != nil {
			pw.CloseWithError(err)
			return
		}
		if _, err := io.Copy(excelPart, excelFile); err != nil {
			pw.CloseWithError(err)
			return
		}

		// zip 文件
		zipFile, err := os.Open(zipPath)
		if err != nil {
			pw.CloseWithError(err)
			return
		}
		defer zipFile.Close()

		zipPart, err := writer.CreateFormFile("zip", filepath.Base(zipPath))
		if err != nil {
			pw.CloseWithError(err)
			return
		}
		if _, err := io.Copy(zipPart, zipFile); err != nil {
			pw.CloseWithError(err)
			return
		}
	}()
	fmt.Println("excel,zip copy done: ", zipPath)
	req, err := http.NewRequestWithContext(context.Background(), "POST", apiURL, pr)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", token)

	client := &http.Client{
		Timeout: 30 * time.Minute, // 根据文件大小调整
	}
	fmt.Println("start http client do")
	now := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("接口返回 %s, body: %s", resp.Status, string(body))
	}
	fmt.Println("time: ", time.Since(now))
	fmt.Println("接口返回:", string(body))
	// 解析 JSON
	var apiResp ApiResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return fmt.Errorf("解析接口返回失败: %v", err)
	}

	if apiResp.Status != 0 {
		// status != 0 表示失败
		return fmt.Errorf("接口返回失败: %s", apiResp.Msg)
	}
	return nil
}
func loadUploaded(file string) map[string]bool {
	uploaded := make(map[string]bool)
	data, err := os.ReadFile(file)
	if err != nil {
		return uploaded
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, " -> ", 2)
		if len(parts) > 0 {
			uploaded[parts[0]] = true
		}
	}
	return uploaded
}
