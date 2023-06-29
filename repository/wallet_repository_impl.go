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
	query := "INSERT INTO wallet (id_customer, balance) VALUES($1, $2) RETURNING id"
	tx.QueryRowContext(ctx, query, request.IdCustomer, request.Balance).Scan(&request.Id)

	wallet = entity.Wallet{
		Id:         request.Id,
		IdCustomer: request.IdCustomer,
		Balance:    request.Balance,
	}

	return
}

func (repository *WalletRepositoryImpl) GetWallet(ctx context.Context, tx *sql.Tx, idCustomer uuid.UUID) (wallet entity.Wallet, err error) {
	query := "SELECT id,balance FROM wallet WHERE id_customer=$1"
	tx.QueryRowContext(ctx, query, idCustomer).Scan(&wallet.Id, &wallet.Balance)
	wallet.IdCustomer = idCustomer

	return
}

func (repository *WalletRepositoryImpl) UpdateWallet(ctx context.Context, tx *sql.Tx, request entity.Wallet) (wallet entity.Wallet, err error) {
	query := "UPDATE wallet SET balance=balance+$1 WHERE id_customer=$2"
	_, err = tx.ExecContext(ctx, query, request.Balance, request.IdCustomer)
	if err != nil {
		return
	}

	wallet = entity.Wallet{
		IdCustomer: request.IdCustomer,
		Balance:    request.Balance,
	}

	return

}
