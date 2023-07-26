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

type CustomerUsecaseImpl struct {
	Repository repository.CustomerRepository
	Validate   *validator.Validate
	Timeout    time.Duration
}

func NewCustomerUsecaseImpl(repository repository.CustomerRepository, validate *validator.Validate, timeout time.Duration) CustomerUsecase {
	return &CustomerUsecaseImpl{
		Repository: repository,
		Validate:   validate,
		Timeout:    timeout,
	}
}

func (usecase *CustomerUsecaseImpl) RegisterCustomer(ctx context.Context, request request.CustomerRegister) (response response.CustomerRegister, err error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.Timeout*time.Second)
	defer cancel()

	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}

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
	customerResponse, err := usecase.Repository.RegisterCustomer(ctx, customer)
	if err != nil {
		return
	}

	response = helper.ToResponseCustomerRegister(customerResponse)

	return

}

func (usecase *CustomerUsecaseImpl) LoginCustomer(ctx context.Context, request request.CustomerLogin) (id uuid.UUID, result bool, err error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.Timeout*time.Second)
	defer cancel()

	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}

	id, passwordHashed, err := usecase.Repository.LoginCustomer(ctx, request.Email)
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
	ctx, cancel := context.WithTimeout(ctx, usecase.Timeout*time.Second)
	defer cancel()

	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}

	customer := entity.Customer{
		Name:      request.Name,
		Email:     request.Email,
		UpdatedAt: int32(time.Now().Unix()),
	}

	customerResponse, err := usecase.Repository.UpdateCustomer(ctx, customer)
	if err != nil {
		return
	}

	response = helper.ToResponseCustomerUpdate(customerResponse)

	return
}

func (usecase *CustomerUsecaseImpl) DeleteCustomer(ctx context.Context, request request.CustomerDelete) (err error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.Timeout*time.Second)
	defer cancel()

	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}

	deletedTime := int32(time.Now().Unix())

	err = usecase.Repository.DeleteCustomer(ctx, request.Email, deletedTime)
	if err != nil {
		return
	}

	return
}
