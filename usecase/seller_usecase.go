package usecase

import (
	"context"

	"github.com/dihanto/go-toko/model/web/request"
	"github.com/dihanto/go-toko/model/web/response"
	"github.com/google/uuid"
)

type SellerUsecase interface {
	RegisterSeller(ctx context.Context, request request.SellerRegister) (response response.SellerRegister, err error)
	LoginSeller(ctx context.Context, request request.SellerLogin) (id uuid.UUID, result bool, err error)
	UpdateSeller(ctx context.Context, request request.SellerUpdate) (response response.SellerUpdate, err error)
	DeleteSeller(ctx context.Context, request request.SellerDelete) (err error)
}
