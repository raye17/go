package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	r := echo.New()
	r.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "test",
			"data":    "data",
		})
	})
	r.Start("127.0.0.1:8899")
}
