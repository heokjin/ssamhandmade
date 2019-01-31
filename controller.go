package main

import (
	// "fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Index struct {
	SingleCatagory []IndexSingleCatagory
}

type IndexSingleCatagory struct {
	Href         string
	ImgSrc       string
	Price        string
	CatagoryName string
}

func Hello(c echo.Context) error {
	index := Index{}
	singleCatagory := make([]IndexSingleCatagory, 2)
	shoes := IndexSingleCatagory{}
	shoes.CatagoryName = "귀여운 신발"
	shoes.Href = "/shop"
	shoes.ImgSrc = "img/ssam-bg-img/1.png"
	shoes.Price = "From $180"

	shoes2 := IndexSingleCatagory{}
	shoes2.CatagoryName = "aaaaaa"
	shoes2.Href = "/shop"
	shoes2.ImgSrc = "img/ssam-bg-img/2.png"
	shoes2.Price = "From $80"

	singleCatagory[0] = shoes
	singleCatagory[1] = shoes2
	index.SingleCatagory = singleCatagory

	return c.Render(http.StatusOK, "index.html", index)
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
