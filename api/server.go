package api

import (
	"github.com/labstack/echo/v4"
	"github.com/timothypattikawa/amole-services/product-service/internal/config"
	"github.com/timothypattikawa/amole-services/product-service/pkg/exception"
	"log"
)

func RunServer(handler func(echo *echo.Echo), config *config.Config) {

	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = exception.CostumeErrorAdvice
	handler(e)

	err := e.Start(":" + config.NewServer().Port)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
