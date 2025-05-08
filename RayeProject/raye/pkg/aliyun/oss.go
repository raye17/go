package aliyun

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"raye/demo/config"
	"strings"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
)

var OssClient *oss.Client

func NewOss() {
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewStaticCredentialsProvider(config.AppConfig.Oss.AccessKeyId, config.AppConfig.Oss.AccessKeySecret)).
		WithRegion(config.AppConfig.Oss.Region)
	client := oss.NewClient(cfg)
	OssClient = client
}
func PutOss(filePath string, mediaType string, needRemove bool) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("文件打开失败")
		return "", err
	}
	defer file.Close()
	if mediaType == "" {
		mediaType = "default/"
	}
	if needRemove {
		defer func() {
			os.Remove(filePath)
		}()
	}
	objectName := mediaType + "/" + filepath.Base(filePath)

	// 根据扩展名动态设置Content-Type
	contentType := "application/octet-stream"
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".gif":
		contentType = "image/gif"
	}
	request := &oss.PutObjectRequest{
		Bucket:      oss.Ptr(config.AppConfig.Oss.Bucket),
		Key:         oss.Ptr(objectName),
		Body:        file,
		ContentType: oss.Ptr(contentType),
	}
	_, err = OssClient.PutObject(context.Background(), request)
	if err != nil {
		log.Println("上传失败")
		fmt.Println(config.AppConfig.Oss.Bucket)
		return "", err
	}
	// 构造图片访问 URL
	imageURL := fmt.Sprintf("https://%s.%s/%s", config.AppConfig.Oss.Bucket, config.AppConfig.Oss.Endpoint, objectName)
	return imageURL, nil
}

func PutOssFromBytes(filePath string, content []byte, mediaType string) (string, error) {
	if mediaType == "" {
		mediaType = "default/"
	}
	objectName := mediaType + "/" + filepath.Base(filePath)

	// 根据扩展名动态设置Content-Type
	contentType := "application/octet-stream"
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".gif":
		contentType = "image/gif"
	}
	request := &oss.PutObjectRequest{
		Bucket:      oss.Ptr(config.AppConfig.Oss.Bucket),
		Key:         oss.Ptr(objectName),
		Body:        bytes.NewReader(content),
		ContentType: oss.Ptr(contentType),
	}
	_, err := OssClient.PutObject(context.Background(), request)
	if err != nil {
		log.Println("上传失败")
		fmt.Println(config.AppConfig.Oss.Bucket)
		return "", err
	}
	// 构造图片访问 URL
	imageURL := fmt.Sprintf("https://%s.%s/%s", config.AppConfig.Oss.Bucket, config.AppConfig.Oss.Endpoint, objectName)
	return imageURL, nil
}
func ListObjects(name string) ([]string, error) {
	if name == "" {
		name = "default/"
	}
	prefix := name + "/"
	var files []string

	request := &oss.ListObjectsV2Request{
		Bucket: oss.Ptr(config.AppConfig.Oss.Bucket),
		Prefix: oss.Ptr(prefix),
	}
	paginator := OssClient.NewListObjectsV2Paginator(request)

	for paginator.HasNext() {
		page, err := paginator.NextPage(context.Background())
		if err != nil {
			return nil, err
		}
		for _, obj := range page.Contents {
			// 拼接完整URL
			url := fmt.Sprintf("https://%s.%s/%s", config.AppConfig.Oss.Bucket, config.AppConfig.Oss.Endpoint, *obj.Key)
			files = append(files, url)
		}
	}
	return files, nil
}
