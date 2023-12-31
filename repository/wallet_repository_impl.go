package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-toko/helper"
	"github.com/dihanto/go-toko/model/entity"
	"github.com/google/uuid"
)

type WalletRepositoryImpl struct {
	Database *sql.DB
}

func NewWalletRepository(database *sql.DB) WalletRepository {
	return &WalletRepositoryImpl{
		Database: database,
	}
}

func (repository *WalletRepositoryImpl) AddWallet(ctx context.Context, request entity.Wallet) (wallet entity.Wallet, err error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	query := "INSERT INTO wallets (id_customer, balance, created_at) VALUES($1, $2, $3) RETURNING id"
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

func (repository *WalletRepositoryImpl) GetWallet(ctx context.Context, idCustomer uuid.UUID) (wallet entity.Wallet, err error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	query := "SELECT id,balance,created_at,updated_at FROM wallets WHERE id_customer=$1"
	err = tx.QueryRowContext(ctx, query, idCustomer).Scan(&wallet.Id, &wallet.Balance, &wallet.CreatedAt, &wallet.UpdatedAt)
	if err != nil {
		return
	}
	wallet.IdCustomer = idCustomer

	return
}

func (repository *WalletRepositoryImpl) UpdateWallet(ctx context.Context, request entity.Wallet) (wallet entity.Wallet, err error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)
	
	query := "UPDATE wallets SET balance=balance+$1, updated_at=$2 WHERE id_customer=$3 RETURNING created_at, balance"
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
