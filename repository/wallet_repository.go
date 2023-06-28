package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-toko/model/entity"
)

type WalletRepository interface {
	AddWallet(ctx context.Context, tx *sql.Tx, request entity.Wallet) (wallet entity.Wallet, err error)
	GetWallet(ctx context.Context, tx *sql.Tx, id int) (wallet entity.Wallet, err error)
	UpdateWallet(ctx context.Context, tx *sql.Tx, request entity.Wallet) (wallet entity.Wallet, err error)
}
