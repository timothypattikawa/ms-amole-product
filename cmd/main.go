package main

import (
	"github.com/labstack/echo/v4"
	"github.com/timothypattikawa/amole-services/product-service/api"
	"github.com/timothypattikawa/amole-services/product-service/internal/config"
	"github.com/timothypattikawa/amole-services/product-service/internal/handler"
	"github.com/timothypattikawa/amole-services/product-service/internal/repository"
	"github.com/timothypattikawa/amole-services/product-service/internal/service"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	v := config.LoadViper(env)
	newConfig := config.NewConfig(v)
	dbConnection := newConfig.NewDatabaseConfig("postgres").GetDbConnection()

	productRepository := repository.NewProductRepository(dbConnection)
	productService := service.NewProductService(v, productRepository)
	productHandler := handler.NewProductHandler(productService)

	api.RunServer(func(e *echo.Echo) {
		handler.Handler(e, productHandler)
	}, newConfig)
}
