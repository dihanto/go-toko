package usecase

import (
	"context"

	"github.com/dihanto/go-toko/model/web/request"
	"github.com/dihanto/go-toko/model/web/response"
	"github.com/google/uuid"
)

type WalletUsecase interface {
	AddWallet(ctx context.Context, request request.AddWallet) (wallet response.AddWallet, err error)
	GetWallet(ctx context.Context, idCustomer uuid.UUID) (wallet response.GetWallet, err error)
	UpdateWallet(ctx context.Context, request request.UpdateWallet) (wallet response.UpdateWallet, err error)
}
