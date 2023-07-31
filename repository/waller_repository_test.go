package repository

import (
	"context"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/dihanto/go-toko/model/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestAddWallet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}

	walletPayload := entity.Wallet{
		Id:         1,
		IdCustomer: uuid.New(),
		Balance:    20000,
		CreatedAt:  int32(time.Now().Unix()),
	}

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO wallets \\(id_customer, balance, created_at\\) VALUES\\(\\$1, \\$2, \\$3\\) RETURNING id").
		WithArgs(walletPayload.IdCustomer, walletPayload.Balance, walletPayload.CreatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	repo := NewWalletRepository(db)
	_, err = repo.AddWallet(context.Background(), walletPayload)
	assert.NoError(t, err)
}

func TestGetWallet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}

	idCustomer := uuid.New()

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id,balance,created_at,updated_at FROM wallets WHERE id_customer=\\$1").
		WithArgs(idCustomer).WillReturnRows(sqlmock.NewRows([]string{"id", "balance", "created_at", "updated_at"}).
		AddRow(1, 10000, "12345678", "12345678"))

	repo := NewWalletRepository(db)
	_, err = repo.GetWallet(context.Background(), idCustomer)
	assert.NoError(t, err)
}

func TestUpdateWallet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}

	walletPayload := entity.Wallet{
		Id:         1,
		Balance:    1000,
		UpdatedAt:  int32(time.Now().Unix()),
		IdCustomer: uuid.New(),
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("UPDATE wallets SET balance=balance+$1, updated_at=$2 WHERE id_customer=$3 RETURNING created_at, balance")).
		WithArgs(walletPayload.Balance, walletPayload.UpdatedAt, walletPayload.IdCustomer).
		WillReturnRows(sqlmock.NewRows([]string{"balance", "updated_at"}).AddRow(1000, "12345678"))

	repo := NewWalletRepository(db)
	_, err = repo.UpdateWallet(context.Background(), walletPayload)
	assert.NoError(t, err)
}
