package main

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("TEST")
	t := &Template{
		templates: template.Must(template.ParseGlob("ssam/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.Static("/", "static")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", Hello)
	e.GET("/shop", Shop)
	e.GET("/product-details", ProductDetails)
	e.GET("/cart", Cart)
	e.GET("/checkout", CheckOut)

	// Start server
	e.Start(":80")
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
