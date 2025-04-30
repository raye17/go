package service

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"raye/demo/pkg/utlis/e"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadImg(c *gin.Context) {
	// 限制上传文件大小
	// 图片最大8MB，视频最大50MB
	var maxSize int64
	if c.Request.Header.Get("Content-Type") == "video/mp4" || c.Request.Header.Get("Content-Type") == "video/avi" || c.Request.Header.Get("Content-Type") == "video/mov" {
		maxSize = 5 << 30
	} else {
		maxSize = 8 << 20
	}
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxSize)

	// 从表单获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取文件失败: " + err.Error()})
		return
	}

	// 验证文件类型
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
		"video/mp4":  true,
		"video/avi":  true,
		"video/mov":  true,
	}
	if !allowedTypes[file.Header.Get("Content-Type")] {
		// 处理分块上传时的文件类型验证
		if strings.HasPrefix(file.Header.Get("Content-Type"), "multipart/form-data") {
			// 分块上传时，文件类型可能被设置为multipart/form-data，这里假设所有分块都是同一类型，且与原始文件一致
			contentType := strings.TrimPrefix(c.Request.Header.Get("Original-Content-Type"), "multipart/form-data; boundary=")
			if !allowedTypes[contentType] {
				c.JSON(http.StatusBadRequest, gin.H{"error": "不支持的文件类型"})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "不支持的文件类型"})
			return
		}
	}

	// 处理分块上传
	chunkNumberStr := c.DefaultQuery("chunkNumber", "1")
	totalChunksStr := c.DefaultQuery("totalChunks", "1")
	chunkNumber, _ := strconv.Atoi(chunkNumberStr)
	totalChunks, _ := strconv.Atoi(totalChunksStr)

	// 创建保存目录
	dir, _ := os.Getwd()
	var destDir string
	if strings.Contains(file.Header.Get("Content-Type"), "video/") {
		destDir = filepath.Join(dir, "/runtime", "videos")
	} else {
		destDir = filepath.Join(dir, "/runtime", "imgs")
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "保存分块文件失败: " + err.Error()})
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
					c.JSON(http.StatusInternalServerError, gin.H{"error": "创建文件失败: " + err.Error()})
					return
				}
				defer outputFile.Close()

				for i := 1; i <= totalChunks; i++ {
					chunkPath := filepath.Join(tempDir, fmt.Sprintf("chunk_%d", i))
					chunkFile, err := os.Open(chunkPath)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "打开分块文件失败: " + err.Error()})
						return
					}
					defer chunkFile.Close()

					_, err = outputFile.ReadFrom(chunkFile)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "合并分块文件失败: " + err.Error()})
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件失败: " + err.Error()})
			return
		}
	}

	// 获取host地址
	host := c.Request.Host
	relativePath, _ := filepath.Rel(filepath.Join(dir, "runtime"), destPath)
	imgUrl := fmt.Sprintf("http://%s/static/%s", host, filepath.ToSlash(relativePath))
	ResponseMsg(c, e.Success, "ok", nil, map[string]string{
		"url": imgUrl,
	})
}
