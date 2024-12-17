package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/timothypattikawa/amole-services/product-service/internal/repository/postgres"
	"time"
)

type ProductRepository interface {
	ExecTx(ctx context.Context, fn func(queries *postgres.Queries) error) error
	Exec(ctx context.Context, fn func(queries *postgres.Queries) error) error
}

type ProductRepositoryImplementation struct {
	db *pgxpool.Pool
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
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	queries := postgres.New(p.db)
	err := fn(queries)
	if err != nil {
		return err
	}
	return err
}

func NewProductRepository(db *pgxpool.Pool) ProductRepository {
	return &ProductRepositoryImplementation{db: db}
}
