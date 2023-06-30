package usecase

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-toko/helper"
	"github.com/dihanto/go-toko/model/entity"
	"github.com/dihanto/go-toko/model/web/request"
	"github.com/dihanto/go-toko/model/web/response"
	"github.com/dihanto/go-toko/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type WalletUsecaseImpl struct {
	Repository repository.WalletRepository
	Db         *sql.DB
	Validate   *validator.Validate
	Timeout    int
}

func NewWalletUsecase(repository repository.WalletRepository, db *sql.DB, validate *validator.Validate, timeout int) WalletUsecase {
	return &WalletUsecaseImpl{
		Repository: repository,
		Db:         db,
		Validate:   validate,
		Timeout:    timeout,
	}
}

func (usecase *WalletUsecaseImpl) AddWallet(ctx context.Context, request request.AddWallet) (wallet response.AddWallet, err error) {
	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}

	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	requestRepo := entity.Wallet{
		IdCustomer: request.IdCustomer,
		Balance:    request.Balance,
	}

	response, err := usecase.Repository.AddWallet(ctx, tx, requestRepo)
	if err != nil {
		return
	}

	wallet = helper.ToResponseAddWallet(response)

	return
}

func (usecase *WalletUsecaseImpl) GetWallet(ctx context.Context, idCustomer uuid.UUID) (wallet response.GetWallet, err error) {
	err = usecase.Validate.Var(idCustomer, "required")
	if err != nil {
		return
	}

	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	response, err := usecase.Repository.GetWallet(ctx, tx, idCustomer)
	if err != nil {
		return
	}

	wallet = helper.ToResponseGetWallet(response)

	return
}

func (usecase *WalletUsecaseImpl) UpdateWallet(ctx context.Context, request request.UpdateWallet) (wallet response.UpdateWallet, err error) {
	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}

	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	requestRepo := entity.Wallet{
		IdCustomer: request.IdCustomer,
		Balance:    request.Balance,
	}

	response, err := usecase.Repository.UpdateWallet(ctx, tx, requestRepo)
	if err != nil {
		return
	}

	wallet = helper.ToResponseUpdateWallet(response)

	return
}
