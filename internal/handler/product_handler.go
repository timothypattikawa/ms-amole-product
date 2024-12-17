package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/timothypattikawa/amole-services/product-service/internal/dto"
	"github.com/timothypattikawa/amole-services/product-service/internal/service"
	"github.com/timothypattikawa/amole-services/product-service/pkg/exception"
	"log"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h ProductHandler) GetProducts(e echo.Context) error {
	log.Printf("Get all products")
	products, err := h.productService.GetAllProducts(e.Request().Context())
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, dto.BaseResponse{
		Data: products,
	})
}

func (h ProductHandler) GetProductById(e echo.Context) error {
	param := e.Param("product_id")
	log.Printf("Get all products by id %s", param)
	productId, err := strconv.Atoi(param)
	if err != nil {
		log.Printf("Error converting %s to int", param)
		return exception.NewInternalServerError("Something went wrong!!")
	}

	product, err := h.productService.GetProductById(e.Request().Context(), int64(productId))
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, dto.BaseResponse{
		Data: product,
	})
}
