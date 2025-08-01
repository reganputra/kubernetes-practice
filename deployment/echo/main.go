package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		hostname, _ := os.Hostname()
		env := os.Getenv("MISSING_ENV")
		if env == "" {
			os.Exit(1)
		}
		return c.String(http.StatusOK, "Hello, World! using v4 on "+hostname+" in "+env+" environment")
	})

	app.Start(":8080")
}
