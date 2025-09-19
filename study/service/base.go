package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Failed = 1
	Ok     = 0

	RetryCode = 2
)

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Code   int         `json:"code"`
	Error  error       `json:"error"`
}

func Success(c *gin.Context, msg string, datas ...interface{}) {
	var data interface{}

	if datas != nil {
		data = datas[0]
	} else {
		data = struct{}{}
	}
	c.JSON(http.StatusOK, Response{
		Status: Ok,
		Code:   Ok,
		Data:   data,
		Msg:    msg,
	})
	c.Abort()
}

// Error 统一错误返回
func Error(c *gin.Context, err error) {

	c.JSON(http.StatusOK, Response{
		Code:   Failed,
		Status: Failed,
		Msg:    err.Error(),
		Data:   struct{}{},
	})

	c.Abort()
}

// 重试
func Retry(c *gin.Context, err error) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	c.JSON(http.StatusOK, Response{
		Code: RetryCode,
		Msg:  errMsg,
		Data: struct{}{},
	})

	c.Abort()
}
