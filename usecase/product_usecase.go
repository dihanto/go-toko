package usecase

import (
	"context"

	"github.com/dihanto/go-toko/model/web/request"
	"github.com/dihanto/go-toko/model/web/response"
)

type ProductUsecase interface {
	AddProduct(ctx context.Context, request request.AddProduct) (product response.AddProduct, err error)
	GetProduct(ctx context.Context) (products []response.GetProduct, err error)
	FindById(ctx context.Context, id int) (product response.FindById, err error)
	UpdateProduct(ctx context.Context, request request.UpdateProduct) (product response.UpdateProduct, err error)
	DeleteProduct(ctx context.Context, id int) (err error)
}
