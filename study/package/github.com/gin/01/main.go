package main

import (
	"log"
	"net/http"
	"time"

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
	r.GET("say", SayHello)
	r.Run(":8080")
}
func SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":   "code~~",
		"status": "status-ok",
		"ip":     c.ClientIP(),
	})
	//return
	log.Println("ssssss")
	time.Sleep(time.Second * 3)
	log.Println("3秒后")
}
