package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	fmt.Println("Hello")
	return c.Render(http.StatusOK, "index.html", "World")
}
