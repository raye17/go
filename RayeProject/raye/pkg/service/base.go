package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"msg"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseMsg(c *gin.Context, status int, msg string, err error, data interface{}) {
	response := Response{
		Status:  status,
		Message: msg,
		Data:    data,
	}
	if err != nil {
		response.Error = err.Error()
	}
	c.JSON(http.StatusOK, response)
}
