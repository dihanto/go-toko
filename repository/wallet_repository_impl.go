package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-toko/model/entity"
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

func (repository *WalletRepositoryImpl) GetWallet(ctx context.Context, tx *sql.Tx, id int) (wallet entity.Wallet, err error) {
	query := "SELECT balance FROM wallet WHERE id=$1"
	tx.QueryRowContext(ctx, query, id).Scan(&wallet.Balance)
	wallet.Id = id

	return
}

func (repository *WalletRepositoryImpl) UpdateWallet(ctx context.Context, tx *sql.Tx, request entity.Wallet) (wallet entity.Wallet, err error) {
	query := "UPDATE wallet SET balance=$1 WHERE id=$2"
	_, err = tx.ExecContext(ctx, query, request.Balance, request.Id)
	if err != nil {
		return
	}

	wallet = entity.Wallet{
		Id:      request.Id,
		Balance: request.Balance,
	}

	return

}
