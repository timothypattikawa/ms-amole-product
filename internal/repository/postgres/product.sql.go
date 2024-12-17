// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: product.sql

package postgres

import (
	"context"
)

const getAllProduct = `-- name: GetAllProduct :many
SELECT tbap_id, tbap_name, tbap_price, tbap_description, tbap_created_at, tbap_updated_at
FROM tb_amole_product
`

func (q *Queries) GetAllProduct(ctx context.Context) ([]TbAmoleProduct, error) {
	rows, err := q.db.Query(ctx, getAllProduct)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TbAmoleProduct
	for rows.Next() {
		var i TbAmoleProduct
		if err := rows.Scan(
			&i.TbapID,
			&i.TbapName,
			&i.TbapPrice,
			&i.TbapDescription,
			&i.TbapCreatedAt,
			&i.TbapUpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductById = `-- name: GetProductById :one
SELECT tbap_id, tbap_name, tbap_price, tbap_description, tbap_created_at, tbap_updated_at
FROM tb_amole_product WHERE tbap_id = $1
`

func (q *Queries) GetProductById(ctx context.Context, tbapID int64) (TbAmoleProduct, error) {
	row := q.db.QueryRow(ctx, getProductById, tbapID)
	var i TbAmoleProduct
	err := row.Scan(
		&i.TbapID,
		&i.TbapName,
		&i.TbapPrice,
		&i.TbapDescription,
		&i.TbapCreatedAt,
		&i.TbapUpdatedAt,
	)
	return i, err
}
