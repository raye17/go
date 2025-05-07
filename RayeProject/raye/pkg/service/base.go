package service

import (
	"net/http"
	"raye/demo/pkg/utlis/e"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func ResponseMsg(c *gin.Context, status int, msg string, err error, data interface{}) {
	response := Response{
		Status: status,
		Msg:    msg,
		Data:   data,
	}
	if err != nil {
		response.Error = err.Error()
	}
	c.JSON(http.StatusOK, response)
}
func NotLoginRes(c *gin.Context, msg string) {

	c.JSON(e.Success, Response{
		Status: 1,
		Msg:    msg,
		Data:   nil,
	})
	c.Abort()
}
