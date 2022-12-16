package main

import (
	"github.com/labstack/echo/v4"
	swagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/swag/*", swagger.WrapHandler)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello,world!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
