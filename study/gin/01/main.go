package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	group = "/grooo"
	p     = "err"
	p1    = "/err1"
)

func main() {
	r := gin.Default()
	r.Group("")
	t := r.Group(group)
	t.GET(p, test)
	t.GET(p1, test)
	fmt.Println(group + p1)
	t.GET("hello", test)
	r.GET("/hello", test)
	r.GET("hello2", test)
	d := r.Group("/dd")
	d.GET("hello", test)
	d.GET("/hello3", test)
	c := r.Group("/cc////")
	c.GET("hello", test)
	c.GET("//hello2", test)
	url := fmt.Sprintf("%s://%s:%d%s", "http", "10.9.8.23", 20001, "prefis/edf/rt")
	fmt.Println(url)
	r.Run(":8899")
}
func test(c *gin.Context) {
	c.JSON(0, "hello,world")
}
