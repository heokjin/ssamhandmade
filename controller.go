package main

import (
	// "fmt"
	"net/http"

	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "World")
}

func Shop(c echo.Context) error {
	return c.Render(http.StatusOK, "shop.html", "World")
}

func ProductDetails(c echo.Context) error {
	return c.Render(http.StatusOK, "product-details.html", "World")
}

func Cart(c echo.Context) error {
	return c.Render(http.StatusOK, "cart.html", "World")
}

func CheckOut(c echo.Context) error {
	return c.Render(http.StatusOK, "checkout.html", "World")
}
