package usecase

import (
	"context"

	"github.com/dihanto/go-toko/model/web/request"
	"github.com/dihanto/go-toko/model/web/response"
)

type OrderUsecase interface {
	AddOrder(ctx context.Context, request request.AddOrder) (order response.AddOrder, err error)
	FindOrder(ctx context.Context, id int) (orderDetail response.FindOrder, err error)
}
