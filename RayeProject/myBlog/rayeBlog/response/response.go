package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, httpStatus int, code int, msg string, data gin.H) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// Success 成功
func Success(c *gin.Context, msg string, data gin.H) {
	Response(c, http.StatusOK, 200, msg, data)
}

// Fail 失败
func Fail(c *gin.Context, msg string, data gin.H) {
	Response(c, http.StatusOK, 400, msg, data)
}
