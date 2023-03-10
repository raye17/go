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
		return c.JSON(http.StatusOK, echo.Map{
			"message": "根目录",
			"data":    "hello",
		})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
