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

type CustomerUsecaseImpl struct {
	Repository repository.CustomerRepository
	Database   *sql.DB
	Validate   *validator.Validate
	Timeout    int
}

func NewCustomerUsecaseImpl(repository repository.CustomerRepository, database *sql.DB, validate *validator.Validate, timeout int) CustomerUsecase {
	return &CustomerUsecaseImpl{
		Repository: repository,
		Database:   database,
		Validate:   validate,
		Timeout:    timeout,
	}
}

func (usecase *CustomerUsecaseImpl) RegisterCustomer(ctx context.Context, request request.CustomerRegister) (response response.CustomerRegister, err error) {
	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}

	tx, err := usecase.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	password, err := helper.HashPassword(request.Password)
	if err != nil {
		return
	}

	customer := entity.Customer{
		Id:           uuid.New(),
		Email:        request.Email,
		Name:         request.Name,
		Password:     password,
		RegisteredAt: int32(time.Now().Unix()),
	}
	customerResponse, err := usecase.Repository.RegisterCustomer(ctx, tx, customer)
	if err != nil {
		return
	}

	response = helper.ToResponseCustomerRegister(customerResponse)

	return

}

func (usecase *CustomerUsecaseImpl) LoginCustomer(ctx context.Context, request request.CustomerLogin) (id uuid.UUID, result bool, err error) {
	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}

	tx, err := usecase.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	id, passwordHashed, err := usecase.Repository.LoginCustomer(ctx, tx, request.Email)
	if err != nil {
		return
	}

	result, err = helper.CheckPasswordHash(passwordHashed, request.Password)

	if !result {
		return
	}
	return

}

func (usecase *CustomerUsecaseImpl) UpdateCustomer(ctx context.Context, request request.CustomerUpdate) (response response.CustomerUpdate, err error) {
	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}

	tx, err := usecase.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	customer := entity.Customer{
		Name:      request.Name,
		Email:     request.Email,
		UpdatedAt: int32(time.Now().Unix()),
	}

	customerResponse, err := usecase.Repository.UpdateCustomer(ctx, tx, customer)
	if err != nil {
		return
	}

	response = helper.ToResponseCustomerUpdate(customerResponse)

	return
}

func (usecase *CustomerUsecaseImpl) DeleteCustomer(ctx context.Context, request request.CustomerDelete) (err error) {
	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}

	tx, err := usecase.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	deletedTime := int32(time.Now().Unix())

	err = usecase.Repository.DeleteCustomer(ctx, tx, request.Email, deletedTime)
	if err != nil {
		return
	}

	return
}
