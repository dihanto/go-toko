package usecase

import (
	"context"

	"github.com/dihanto/go-toko/model/web/request"
	"github.com/dihanto/go-toko/model/web/response"
	"github.com/google/uuid"
)

type CustomerUsecase interface {
	RegisterCustomer(ctx context.Context, request request.CustomerRegister) (response response.CustomerRegister, err error)
	LoginCustomer(ctx context.Context, request request.CustomerLogin) (id uuid.UUID, result bool, err error)
	UpdateCustomer(ctx context.Context, request request.CustomerUpdate) (response response.CustomerUpdate, err error)
	DeleteCustomer(ctx context.Context, request request.CustomerDelete) (err error)
}
