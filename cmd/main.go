package main

import (
	"log"
	"net"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/timothypattikawa/amole-services/product-service/api"
	pb "github.com/timothypattikawa/amole-services/product-service/api/grpc/protos/product"
	rpcServer "github.com/timothypattikawa/amole-services/product-service/api/grpc/server"
	"github.com/timothypattikawa/amole-services/product-service/internal/config"
	"github.com/timothypattikawa/amole-services/product-service/internal/handler"
	"github.com/timothypattikawa/amole-services/product-service/internal/repository"
	"github.com/timothypattikawa/amole-services/product-service/internal/repository/rd"
	"github.com/timothypattikawa/amole-services/product-service/internal/service"
	"google.golang.org/grpc"
)

func main() {
	env := os.Getenv("ENV")
	v := config.LoadViper(env)
	newConfig := config.NewConfig(v)
	dbConnection := newConfig.NewDatabaseConfig("postgres").GetDbConnection()

	redisClient := rd.NewRedisConfig(v).GetClient()

	productRepository := repository.NewProductRepository(dbConnection, redisClient)

	productService := service.NewProductService(v, productRepository)
	productHandler := handler.NewProductHandler(productService)

	listen, err := net.Listen("tcp", ":"+v.GetString("service.port-grpc"))
	if err != nil {
		log.Fatalf("error when listen for rpc err::{%v}", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductStockServer(grpcServer, rpcServer.NewServerProductRPC(productRepository))

	grpcServer.Serve(listen)

	api.RunServer(func(e *echo.Echo) {
		handler.Handler(e, productHandler)
	}, newConfig)
}
