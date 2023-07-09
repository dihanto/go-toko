package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-toko/model/entity"
	"github.com/google/uuid"
)

type WalletRepositoryImpl struct {
}

func NewWalletRepository() WalletRepository {
	return &WalletRepositoryImpl{}
}

func (repository *WalletRepositoryImpl) AddWallet(ctx context.Context, tx *sql.Tx, request entity.Wallet) (wallet entity.Wallet, err error) {
	query := "INSERT INTO wallet (id_customer, balance, created_at) VALUES($1, $2, $3) RETURNING id"
	err = tx.QueryRowContext(ctx, query, request.IdCustomer, request.Balance, request.CreatedAt).Scan(&request.Id)
	if err != nil {
		return
	}

	wallet = entity.Wallet{
		Id:         request.Id,
		IdCustomer: request.IdCustomer,
		Balance:    request.Balance,
		CreatedAt:  request.CreatedAt,
	}

	return
}

func (repository *WalletRepositoryImpl) GetWallet(ctx context.Context, tx *sql.Tx, idCustomer uuid.UUID) (wallet entity.Wallet, err error) {
	query := "SELECT id,balance,created_at,updated_at FROM wallet WHERE id_customer=$1"
	err = tx.QueryRowContext(ctx, query, idCustomer).Scan(&wallet.Id, &wallet.Balance, &wallet.CreatedAt, &wallet.UpdatedAt)
	if err != nil {
		return
	}
	wallet.IdCustomer = idCustomer

	return
}

func (repository *WalletRepositoryImpl) UpdateWallet(ctx context.Context, tx *sql.Tx, request entity.Wallet) (wallet entity.Wallet, err error) {
	query := "UPDATE wallet SET balance=balance+$1, updated_at=$2 WHERE id_customer=$3 RETURNING created_at, balance"
	err = tx.QueryRowContext(ctx, query, request.Balance, request.UpdatedAt, request.IdCustomer).Scan(&request.CreatedAt, &request.Balance)
	if err != nil {
		return
	}
	wallet = entity.Wallet{
		IdCustomer: request.IdCustomer,
		Balance:    request.Balance,
		CreatedAt:  request.CreatedAt,
		UpdatedAt:  request.UpdatedAt,
	}

	return

}
