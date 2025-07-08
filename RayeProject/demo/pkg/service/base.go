package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  msg,
		Data: data,
	})
}
func Error(c *gin.Context, msg string, err ...error) {
	var errMsg string
	for _, e := range err {
		errMsg += e.Error() + "\n"
	}
	c.JSON(http.StatusOK, Response{
		Code:  1,
		Msg:   msg,
		Error: errMsg,
		Data:  nil,
	})
	c.Abort()
}
