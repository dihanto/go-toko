package repository

import (
	"context"

	"github.com/dihanto/go-toko/model/entity"
	"github.com/google/uuid"
)

type SellerRepository interface {
	RegisterSeller(ctx context.Context, request entity.Seller) (seller entity.Seller, err error)
	LoginSeller(ctx context.Context, email string) (id uuid.UUID, password string, err error)
	UpdateSeller(ctx context.Context, request entity.Seller) (seller entity.Seller, err error)
	DeleteSeller(ctx context.Context, deleteTime int32, email string) (err error)
}
