package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/gommon/log"
	"github.com/timothypattikawa/amole-services/product-service/internal/repository/postgres"
)

type ProductRepository interface {
	ExecTx(ctx context.Context, fn func(queries *postgres.Queries) error) error
	Exec(ctx context.Context, fn func(queries *postgres.Queries) error) error
	GetProductStock(productId int64) (int64, error)
	UpdateProductStock(productId, stock int64) error
}

type ProductRepositoryImplementation struct {
	db  *pgxpool.Pool
	rdc *redis.Client
}

func (p ProductRepositoryImplementation) ExecTx(ctx context.Context, fn func(queries *postgres.Queries) error) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	begin, err := p.db.Begin(ctx)
	if err != nil {
		return err
	}
	tx := postgres.New(p.db).WithTx(begin)
	err = fn(tx)
	if err != nil {
		begin.Rollback(ctx)
		return err
	}
	err = begin.Commit(ctx)
	if err != nil {
		return err
	}
	return err
}

func (p ProductRepositoryImplementation) Exec(ctx context.Context, fn func(queries *postgres.Queries) error) error {

	queries := postgres.New(p.db)
	err := fn(queries)
	if err != nil {
		return err
	}
	return err
}

func (p ProductRepositoryImplementation) GetProductStock(productId int64) (int64, error) {
	query := fmt.Sprintf("amole|stock|%v", productId)
	rdCmd, err := p.rdc.Get(query).Int64()
	if err != nil {
		log.Errorf("fail get redis key for GetProductStock(%v) err {%v}", query, err)
		return 0, err
	}

	return rdCmd, err
}

func (p ProductRepositoryImplementation) UpdateProductStock(productId, stock int64) error {
	err := p.rdc.Set(fmt.Sprintf("amole|stock|%v", productId), stock, -1).Err()
	return err
}

func NewProductRepository(db *pgxpool.Pool, rdc *redis.Client) ProductRepository {
	return &ProductRepositoryImplementation{
		db:  db,
		rdc: rdc,
	}
}
