package service

import (
	"context"
	"github.com/spf13/viper"
	"github.com/timothypattikawa/amole-services/product-service/internal/repository"
	"github.com/timothypattikawa/amole-services/product-service/internal/repository/postgres"
	"github.com/timothypattikawa/amole-services/product-service/pkg/exception"
	"log"
)

type ProductService interface {
	GetAllProducts(ctx context.Context) (*[]postgres.TbAmoleProduct, error)
	GetProductById(ctx context.Context, id int64) (*postgres.TbAmoleProduct, error)
}

type ProductServiceImpl struct {
	v  *viper.Viper
	pr repository.ProductRepository
}

func (p ProductServiceImpl) GetAllProducts(ctx context.Context) (*[]postgres.TbAmoleProduct, error) {
	var products []postgres.TbAmoleProduct
	err := p.pr.Exec(ctx, func(queries *postgres.Queries) error {
		productDb, err := queries.GetAllProduct(ctx)
		if err != nil {
			log.Printf("failed to get all products: %v", err)
			return exception.NewInternalServerError("something went wrong!!")
		}
		products = append(products, productDb...)
		return nil
	})

	if err != nil {
		log.Printf("failed to get all products: %v", err)
		return nil, err
	}

	return &products, nil
}

func (p ProductServiceImpl) GetProductById(ctx context.Context, id int64) (*postgres.TbAmoleProduct, error) {
	var productResult postgres.TbAmoleProduct
	err := p.pr.Exec(ctx, func(queries *postgres.Queries) error {
		product, err := queries.GetProductById(ctx, id)
		if err != nil {
			log.Printf("failed to get product: %v", err)
			return exception.NewNotFoundError("something went wrong!!")
		}
		productResult = product
		return nil
	})

	if err != nil {
		log.Printf("failed to get product: %v", err)
		return nil, err
	}

	return &productResult, nil
}

func NewProductService(v *viper.Viper, pr repository.ProductRepository) ProductService {
	return &ProductServiceImpl{
		v:  v,
		pr: pr,
	}
}
