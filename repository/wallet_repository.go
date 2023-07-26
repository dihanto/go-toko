package repository

import (
	"context"

	"github.com/dihanto/go-toko/model/entity"
	"github.com/google/uuid"
)

type WalletRepository interface {
	AddWallet(ctx context.Context, request entity.Wallet) (wallet entity.Wallet, err error)
	GetWallet(ctx context.Context, id uuid.UUID) (wallet entity.Wallet, err error)
	UpdateWallet(ctx context.Context, request entity.Wallet) (wallet entity.Wallet, err error)
}
