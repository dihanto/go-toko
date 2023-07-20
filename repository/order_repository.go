package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-toko/model/entity"
)

type OrderRepository interface {
	AddOrder(ctx context.Context, tx *sql.Tx, orderRequest entity.Order, orderDetailRequest entity.OrderDetail) (order entity.Order, orderDetail entity.OrderDetail, err error)
	FindOrder(ctx context.Context, tx *sql.Tx, id int) (order entity.Order, orderDetail entity.OrderDetail, product entity.Product, customerName string, err error)
}
