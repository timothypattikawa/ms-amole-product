package handler

import "github.com/labstack/echo/v4"

func Handler(e *echo.Echo, handler *ProductHandler) {
	e.GET("/v1/products", handler.GetProducts)
	e.GET("/v1/product/:product_id", handler.GetProductById)
}
