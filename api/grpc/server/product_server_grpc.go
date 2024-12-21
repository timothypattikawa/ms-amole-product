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
	if stock < req.QtyStock {
		log.Printf("error cause out of stock product{%v} stock{%v}", req.Id, req.QtyStock)
		return nil, status.New(codes.NotFound, "out of stock").Err()
	} else {
		err = psg.pr.UpdateProductStock(req.Id, stock-req.QtyStock)
		if err != nil {
			return nil, status.New(codes.NotFound, "out of stock").Err()
		}

		return &pb.TakeStockForATCResponse{
			SuccessTakeStock: true,
			Id:               product.TbapID,
			ProductName:      product.TbapName,
			Price:            int64(product.TbapPrice),
		}, nil
	}
}

func (psg *ProductServerGRPC) ProductInfo(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
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

func NewServerProductRPC(pr repository.ProductRepository) *ProductServerGRPC {
	return &ProductServerGRPC{
		pr: pr,
	}
}
