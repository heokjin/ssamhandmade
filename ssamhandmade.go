package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("TEST")
	e := echo.New()
	e.Static("/", "amado")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//	e.GET("/", testa)

	// Start server
	e.Start(":80")
}
