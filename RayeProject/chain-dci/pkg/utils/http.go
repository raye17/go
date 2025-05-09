package utils

import (
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

func Post(url string, jsonStr []byte) (statusCode int, result string) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	statusCode = resp.StatusCode
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("post url:", url)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", string(body))
	result = string(body)
	zap.L().Info("post", zap.Any("url", url), zap.Any("jsonStr", jsonStr), zap.Any("result", result))
	return
}

func Get(url string) (statusCode int, result string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	statusCode = resp.StatusCode
	fmt.Println("response StatusCode:", resp.StatusCode)
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	result = string(body)
	fmt.Println("response Body:", string(body))
	zap.L().Info("Get", zap.Any("url", url), zap.Any("result", result))
	return
}

func PutFromFileUrlWithStream(url, fileName, fileUrl string) (statusCode int, result string) {
	file, err := http.Get(fileUrl)
	if err != nil {
		panic(err)
	}
	defer file.Body.Close()
	fileBody, _ := ioutil.ReadAll(file.Body)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(fileBody))
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("x-oss-meta-rawfilename", fileName)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//panic(err)
		return 400, "执行文件上传失败"
	}
	defer resp.Body.Close()
	statusCode = resp.StatusCode
	body, _ := ioutil.ReadAll(resp.Body)
	result = string(body)
	fmt.Println("put url:", url)
	fmt.Println("fileName :", fileName)
	fmt.Println("response Headers:", resp.Header)
	//fmt.Println("response Body:", string(body))
	fmt.Println("response StatusCode:", statusCode)
	//zap.L().Info("post", zap.Any("url", url), zap.Any("jsonStr", bytes.NewBuffer(fileBody).String()), zap.Any("result", result))
	return
}
