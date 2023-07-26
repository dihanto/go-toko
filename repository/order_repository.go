package repository

import (
	"context"

	"github.com/dihanto/go-toko/model/entity"
)

type OrderRepository interface {
	AddOrder(ctx context.Context, orderRequest entity.Order, orderDetailRequest entity.OrderDetail) (order entity.Order, orderDetail entity.OrderDetail, err error)
	FindOrder(ctx context.Context, id int) (order entity.Order, orderDetail entity.OrderDetail, product entity.Product, customerName string, err error)
}
