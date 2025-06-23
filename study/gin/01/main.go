package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", test)
	r.Run(":8899")
}
func test(c *gin.Context) {
	c.JSON(0, "hello,world")
}
