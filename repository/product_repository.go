package repository

import (
	"context"

	"github.com/dihanto/go-toko/model/entity"
	"github.com/google/uuid"
)

type ProductRepository interface {
	AddProduct(ctx context.Context, request entity.Product) (product entity.Product, err error)
	GetProduct(ctx context.Context) (products []entity.Product, err error)
	FindById(ctx context.Context, id int) (product entity.Product, err error)
	UpdateProduct(ctx context.Context, request entity.Product) (product entity.Product, err error)
	DeleteProduct(ctx context.Context, deleteTime int32, id int) (err error)
	Search(ctx context.Context, search string, offset int, limit int) (products []entity.Product, count string, err error)
	AddProductToWishlist(ctx context.Context, productId int, customerId uuid.UUID) (product entity.Product, err error)
	DeleteProductFromWishlist(ctx context.Context, productId int, customerId uuid.UUID) (product entity.Product, err error)
}
