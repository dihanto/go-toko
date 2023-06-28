package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-toko/model/entity"
)

type OrderRepository interface {
	AddOrder(ctx context.Context, tx *sql.Tx, request entity.Order) (order entity.Order, err error)
	FindOrder(ctx context.Context, tx *sql.Tx, id int) (order entity.Order, product entity.Product, customerName string, err error)
}
