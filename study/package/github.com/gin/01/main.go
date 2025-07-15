package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello,sxy!")
	})
	r.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":   "code~~",
			"status": "status-ok",
			"ip":     c.ClientIP(),
		})
	})
	r.Run(":8080")
}
