package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func main() {
	s := Server{}
	r := gin.Default()
	s.router = r
	r.GET("/", test)
	go func() {
		err := s.listenAndServe(9912)
		if err != nil {
			panic(err)
		}
	}()
	err := s.listenAndServe(9911)
	if err != nil {
		panic(err)
	}
}
func (s *Server) listenAndServe(port int) error {
	return s.router.Run(fmt.Sprintf(":%d", port))
}
func test(c *gin.Context) {
	c.String(200, "hello,sxy!")
}
