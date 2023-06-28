package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/dihanto/go-toko/helper"
	"github.com/dihanto/go-toko/model/entity"
	"github.com/dihanto/go-toko/model/web/request"
	"github.com/dihanto/go-toko/model/web/response"
	"github.com/dihanto/go-toko/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type SellerUsecaseImpl struct {
	Repository repository.SellerRepository
	Db         *sql.DB
	Validate   *validator.Validate
	Timeout    int
}

func NewSellerUsecaseImpl(repository repository.SellerRepository, db *sql.DB, validate *validator.Validate, timeout int) SellerUsecase {
	return &SellerUsecaseImpl{
		Repository: repository,
		Db:         db,
		Validate:   validate,
		Timeout:    timeout,
	}
}

func (usecase *SellerUsecaseImpl) RegisterSeller(ctx context.Context, request request.SellerRegister) (response response.SellerRegister, err error) {
	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	password, err := helper.HashPassword(request.Password)
	if err != nil {
		return
	}

	seller := entity.Seller{
		Id:           uuid.New(),
		Email:        request.Email,
		Name:         request.Name,
		Password:     password,
		RegisteredAt: int32(time.Now().Unix()),
	}

	sellerResponse, err := usecase.Repository.RegisterSeller(ctx, tx, seller)
	if err != nil {
		return
	}

	response = helper.ToResponseSellerRegister(sellerResponse)

	return
}

func (usecase *SellerUsecaseImpl) LoginSeller(ctx context.Context, request request.SellerLogin) (id uuid.UUID, result bool, err error) {
	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	id, password, err := usecase.Repository.LoginSeller(ctx, tx, request.Email)
	if err != nil {
		return
	}

	result, err = helper.CheckPasswordHash(password, request.Password)
	if err != nil {
		return
	}

	return

}

func (usecase *SellerUsecaseImpl) UpdateSeller(ctx context.Context, request request.SellerUpdate) (response response.SellerUpdate, err error) {
	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	seller := entity.Seller{
		Name:      request.Name,
		Email:     request.Email,
		UpdatedAt: int32(time.Now().Unix()),
	}

	sellerResponse, err := usecase.Repository.UpdateSeller(ctx, tx, seller)
	if err != nil {
		return
	}

	response = helper.ToResponseSellerUpdate(sellerResponse)

	return
}

func (usecase *SellerUsecaseImpl) DeleteSeller(ctx context.Context, request request.SellerDelete) (err error) {
	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	deleteTime := int32(time.Now().Unix())

	err = usecase.Repository.DeleteSeller(ctx, tx, deleteTime, request.Email)
	if err != nil {
		return
	}

	return
}
