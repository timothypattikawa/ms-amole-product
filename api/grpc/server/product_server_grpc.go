package server

import (
	"context"
	"log"

	pb "github.com/timothypattikawa/amole-services/product-service/api/grpc/protos/product"
	repository "github.com/timothypattikawa/amole-services/product-service/internal/repository"
	"github.com/timothypattikawa/amole-services/product-service/internal/repository/postgres"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductServerGRPC struct {
	pb.UnimplementedProductStockServer
	pr repository.ProductRepository
}

func (psg *ProductServerGRPC) TakeStockForATC(ctx context.Context, req *pb.TakeStockForATCkRequest) (*pb.TakeStockForATCResponse, error) {
	log.Printf("incoming request to TakeStockForATC req{%v}", req)
	var product postgres.TbAmoleProduct
	err := psg.pr.Exec(ctx, func(q *postgres.Queries) error {
		result, err := q.GetProductById(ctx, req.GetId())
		if err != nil {
			return status.New(codes.NotFound, "invalid argument").Err()
		}

		product = result
		return nil
	})
	if err != nil {
		return nil, status.New(codes.NotFound, "out of stock").Err()
	}

	stock, err := psg.pr.GetProductStock(req.GetId())
	if err != nil {
		return nil, err
	}

	err = psg.pr.UpdateProductStock(req.Id, req.UserCartStockQty)
	if err != nil {
		return nil, status.New(codes.NotFound, "out of stock").Err()
	}

	if stock < req.QtyStock {
		log.Printf("error cause out of stock product{%v} stock{%v}", req.Id, req.QtyStock)
		return nil, status.New(codes.NotFound, "out of stock").Err()
	} else {
		err = psg.pr.UpdateProductStock(req.Id, -req.QtyStock)
		if err != nil {
			return nil, status.New(codes.NotFound, "out of stock").Err()
		}

		log.Println("Success to update product stock")
		return &pb.TakeStockForATCResponse{
			SuccessTakeStock: true,
			Id:               product.TbapID,
			ProductName:      product.TbapName,
			Price:            int64(product.TbapPrice),
		}, nil
	}
}

func (psg *ProductServerGRPC) ProductInfo(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	log.Printf("incoming request to product Info req{%v}", req)
	var product postgres.TbAmoleProduct
	err := psg.pr.Exec(ctx, func(q *postgres.Queries) error {
		result, err := q.GetProductById(ctx, req.GetTbapID())
		if err != nil {
			return status.New(codes.NotFound, "invalid argument").Err()
		}

		product = result
		return nil
	})

	return &pb.ProductResponse{
		TbapID:          product.TbapID,
		TbapName:        product.TbapName,
		TbapPrice:       product.TbapPrice,
		TbapDescription: product.TbapDescription,
	}, err
}

func (psg *ProductServerGRPC) PutBackStock(ctx context.Context, req *pb.PutStockkRequest) (*pb.PutStockResponse, error) {
	err := psg.pr.UpdateProductStock(req.Id, req.QtyStock)
	if err != nil {
		log.Printf("fail put back product stock err{%v}", err)
		return nil, status.New(codes.Canceled, "fail put back product stock").Err()
	}
	return &pb.PutStockResponse{
		SuccessTakeStock: true,
	}, nil
}

func NewServerProductRPC(pr repository.ProductRepository) *ProductServerGRPC {
	return &ProductServerGRPC{
		pr: pr,
	}
}
