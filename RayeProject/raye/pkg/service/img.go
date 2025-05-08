package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"raye/demo/pkg/aliyun"
	"raye/demo/pkg/utlis/e"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileMeta struct {
	Type int                   `json:"type"` // 1:图片 2:视频 3:文件 0:default
	Mask string                `json:"mask"` // 备注(可选)
	File *multipart.FileHeader `json:"-"`    // 上传的文件数据
}

func UploadImg(c *gin.Context) {
	var fileMeta FileMeta
	if err := c.ShouldBind(&fileMeta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析文件信息失败: " + err.Error()})
		return
	}

	// 限制上传文件大小
	// 图片最大80MB，视频最大50MB
	var maxSize int64
	switch fileMeta.Type {
	case 1: // 图片
		maxSize = 10 << 20 // 10MB
	case 2: // 视频
		maxSize = 100 << 20 // 100MB
	default: // 其他文件
		maxSize = 20 << 20 // 20MB
	}
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxSize)

	// 从表单获取文件
	file, err := c.FormFile("file")
	if err != nil {
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
		return
	}

	// 根据文件扩展名自动判断类型
	ext := filepath.Ext(file.Filename)
	fmt.Printf("上传文件: %s, 扩展名: %s\n", file.Filename, ext)

	// 自动设置文件类型
	if fileMeta.Type == 0 {
		switch strings.ToLower(ext) {
		case ".jpg", ".jpeg", ".png", ".gif":
			fileMeta.Type = 1
		case ".mp4", ".avi", ".mov", ".wmv":
			fileMeta.Type = 2
		default:
			fileMeta.Type = 3
		}
		fmt.Printf("自动设置文件类型为: %d\n", fileMeta.Type)
	}

	// 验证文件类型
	allowedTypes := map[int]bool{
		1: true,
		2: true,
		3: true,
		0: true,
	}
	if !allowedTypes[fileMeta.Type] {
		ResponseMsg(c, e.Failed, "不支持的文件类型", nil, nil)
		return
	}

	// 自动分块逻辑 - 每块5MB
	chunkSize := int64(50 << 20) // 5MB
	fileSize := file.Size
	totalChunks := int(fileSize/chunkSize) + 1
	chunkNumber := 1 // 默认第一块

	// 创建保存目录
	dir, _ := os.Getwd()
	var destDir string
	switch fileMeta.Type {
	case 1:
		destDir = filepath.Join(dir, "/runtime", "imgs")
	case 2:
		destDir = filepath.Join(dir, "/runtime", "videos")
	case 3:
		destDir = filepath.Join(dir, "/runtime", "files")
	default:
		destDir = filepath.Join(dir, "/runtime", "default")
	}
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		os.MkdirAll(destDir, 0755)
	}

	// 生成唯一文件名
	fileName := uuid.New().String() + filepath.Ext(file.Filename)
	destPath := filepath.Join(destDir, fileName)

	// 处理分块上传
	if totalChunks > 1 {
		// 创建临时目录
		tempDir := filepath.Join(destDir, "temp", fileName)
		if _, err := os.Stat(tempDir); os.IsNotExist(err) {
			os.MkdirAll(tempDir, 0755)
		}

		// 保存分块文件
		chunkPath := filepath.Join(tempDir, fmt.Sprintf("chunk_%d", chunkNumber))
		if err := c.SaveUploadedFile(file, chunkPath); err != nil {
			ResponseMsg(c, e.Failed, "保存分块文件失败: ", err, nil)
			return
		}

		// 检查是否所有分块都已上传完成
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			if chunkNumber == totalChunks {
				// 合并分块文件
				outputFile, err := os.Create(destPath)
				if err != nil {
					ResponseMsg(c, e.Failed, "创建文件失败: ", err, nil)
					return
				}
				defer outputFile.Close()

				for i := 1; i <= totalChunks; i++ {
					chunkPath := filepath.Join(tempDir, fmt.Sprintf("chunk_%d", i))
					chunkFile, err := os.Open(chunkPath)
					if err != nil {
						ResponseMsg(c, e.Failed, "打开分块文件失败: ", err, nil)
						return
					}
					defer chunkFile.Close()

					_, err = outputFile.ReadFrom(chunkFile)
					if err != nil {
						ResponseMsg(c, e.Failed, "合并分块文件失败: ", err, nil)
						return
					}

					// 删除已合并的分块文件
					os.Remove(chunkPath)
				}

				// 删除临时目录
				os.RemoveAll(tempDir)
			}
		}()

		wg.Wait()
	} else {
		// 普通上传
		if err := c.SaveUploadedFile(file, destPath); err != nil {
			ResponseMsg(c, e.Failed, "保存文件失败: "+err.Error(), err, nil)
			return
		}
	}

	// 获取host地址
	host := c.Request.Host
	relativePath, _ := filepath.Rel(filepath.Join(dir, "runtime"), destPath)
	imgUrl := fmt.Sprintf("http://%s/static/%s", host, filepath.ToSlash(relativePath))
	ResponseMsg(c, e.Success, "ok", nil, map[string]interface{}{
		"url":  imgUrl,
		"meta": fileMeta,
	})
}

func GetUploadedFiles(c *gin.Context) {
	var fileMeta FileMeta
	if err := c.ShouldBind(&fileMeta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析文件信息失败: " + err.Error()})
		return
	}
	fmt.Println("list: ", fileMeta)
	var destDir string
	dir, _ := os.Getwd()
	switch fileMeta.Type {
	case 1:
		destDir = filepath.Join(dir, "/runtime", "imgs")
	case 2:
		destDir = filepath.Join(dir, "/runtime", "videos")
	case 3:
		destDir = filepath.Join(dir, "/runtime", "files")
	default:
		ResponseMsg(c, e.Failed, "无此类型", nil, nil)
	}

	// 检查目录是否存在或为空
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		ResponseMsg(c, e.Success, "ok", nil, []map[string]interface{}{})
		return
	}

	files, err := os.ReadDir(destDir)
	if err != nil {
		ResponseMsg(c, e.Failed, "读取目录失败: "+err.Error(), err, nil)
		return
	}

	// 如果目录为空，返回空数组
	if len(files) == 0 {
		ResponseMsg(c, e.Success, "ok", nil, []map[string]interface{}{})
		return
	}

	var fileUrls []map[string]interface{}
	host := c.Request.Host
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := filepath.Join(destDir, file.Name())
		relativePath, _ := filepath.Rel(filepath.Join(dir, "runtime"), filePath)
		fileUrl := fmt.Sprintf("http://%s/static/%s", host, filepath.ToSlash(relativePath))
		fileUrls = append(fileUrls, map[string]interface{}{
			"url": fileUrl,
		})
	}

	ResponseMsg(c, e.Success, "ok", nil, fileUrls)
}
func PutObject(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		ResponseMsg(c, e.Failed, "解析文件信息失败: "+err.Error(), err, nil)
		return
	}
	imgType := c.PostForm("type")
	ext := filepath.Ext(file.Filename)
	fileName := uuid.New().String() + ext
	dir, _ := os.Getwd()
	destPath := filepath.Join(dir, "runtime", "imgs", fileName)
	if err = c.SaveUploadedFile(file, destPath); err != nil {
		ResponseMsg(c, e.Failed, "保存文件失败: "+err.Error(), err, nil)
		return
	}
	url, err := aliyun.PutOss(destPath, imgType, true)
	if err != nil {
		ResponseMsg(c, e.Failed, "上传失败: "+err.Error(), err, nil)
		return
	}
	ResponseMsg(c, e.Success, "ok", nil, map[string]interface{}{
		"url": url,
	})
}
func PutObjectByBytes(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		ResponseMsg(c, e.Failed, "获取文件信息失败: "+err.Error(), err, nil)
		return
	}
	fmt.Println(file.Filename)
	f, err := file.Open()
	if err != nil {
		ResponseMsg(c, e.Failed, "文件打开失败: "+err.Error(), err, nil)
		return
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		ResponseMsg(c, e.Failed, "读取文件内容失败: "+err.Error(), err, nil)
		return
	}

	imgType := c.PostForm("type")
	ext := filepath.Ext(file.Filename)
	fileName := uuid.New().String() + ext

	url, err := aliyun.PutOssFromBytes(fileName, buf, imgType)
	if err != nil {
		ResponseMsg(c, e.Failed, "上传失败: "+err.Error(), err, nil)
		return
	}
	ResponseMsg(c, e.Success, "ok", nil, map[string]interface{}{
		"url": url,
	})
}
func ListObjects(c *gin.Context) {
	var req struct {
		Name string `form:"name"`
	}
	if err := c.ShouldBind(&req); err != nil {
		ResponseMsg(c, e.Failed, "解析请求失败: "+err.Error(), err, nil)
		return
	}
	fmt.Println("list: ", req)
	url, err := aliyun.ListObjects(req.Name)
	if err != nil {
		ResponseMsg(c, e.Failed, "获取失败: "+err.Error(), err, nil)
		return
	}
	var data []map[string]interface{}
	for _, v := range url {
		data = append(data, map[string]interface{}{
			"url": v,
		})
	}
	ResponseMsg(c, e.Success, "ok", nil, data)
}
