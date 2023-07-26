package usecase

import (
	"context"
	"time"

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
	Validate   *validator.Validate
	Timeout    time.Duration
}

func NewWalletUsecase(repository repository.WalletRepository, validate *validator.Validate, timeout time.Duration) WalletUsecase {
	return &WalletUsecaseImpl{
		Repository: repository,
		Validate:   validate,
		Timeout:    timeout,
	}
}

func (usecase *WalletUsecaseImpl) AddWallet(ctx context.Context, request request.AddWallet) (wallet response.AddWallet, err error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.Timeout*time.Second)
	defer cancel()

	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}

	requestRepo := entity.Wallet{
		IdCustomer: request.IdCustomer,
		Balance:    request.Balance,
		CreatedAt:  int32(time.Now().Unix()),
	}

	response, err := usecase.Repository.AddWallet(ctx, requestRepo)
	if err != nil {
		return
	}

	wallet = helper.ToResponseAddWallet(response)

	return
}

func (usecase *WalletUsecaseImpl) GetWallet(ctx context.Context, idCustomer uuid.UUID) (wallet response.GetWallet, err error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.Timeout*time.Second)
	defer cancel()

	err = usecase.Validate.Var(idCustomer, "required")
	if err != nil {
		return
	}

	response, err := usecase.Repository.GetWallet(ctx, idCustomer)
	if err != nil {
		return
	}

	wallet = helper.ToResponseGetWallet(response)

	return
}

func (usecase *WalletUsecaseImpl) UpdateWallet(ctx context.Context, request request.UpdateWallet) (wallet response.UpdateWallet, err error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.Timeout*time.Second)
	defer cancel()

	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}

	requestRepo := entity.Wallet{
		IdCustomer: request.IdCustomer,
		Balance:    request.Balance,
		UpdatedAt:  int32(time.Now().Unix()),
	}

	response, err := usecase.Repository.UpdateWallet(ctx, requestRepo)
	if err != nil {
		return
	}

	wallet = helper.ToResponseUpdateWallet(response)

	return
}
