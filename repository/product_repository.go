package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-toko/model/entity"
)

type ProductRepository interface {
	AddProduct(ctx context.Context, tx *sql.Tx, request entity.Product) (product entity.Product, err error)
	GetProduct(ctx context.Context, tx *sql.Tx) (products []entity.Product, err error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (product entity.Product, err error)
	UpdateProduct(ctx context.Context, tx *sql.Tx, request entity.Product) (product entity.Product, err error)
	DeleteProduct(ctx context.Context, tx *sql.Tx, deleteTime int32, id int) (err error)
}
