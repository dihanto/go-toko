package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-toko/helper"
	"github.com/dihanto/go-toko/model/entity"
	"github.com/google/uuid"
)

type CustomerRepositoryImpl struct {
	Database *sql.DB
}

func NewCustomerRepositoryImpl(database *sql.DB) CustomerRepository {
	return &CustomerRepositoryImpl{
		Database: database,
	}
}

func (repository *CustomerRepositoryImpl) RegisterCustomer(ctx context.Context, request entity.Customer) (customer entity.Customer, err error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	query := "INSERT INTO customers (id, email, name, password, registered_at) VALUES ($1, $2, $3, $4, $5);"
	_, err = tx.ExecContext(ctx, query, request.Id, request.Email, request.Name, request.Password, request.RegisteredAt)
	if err != nil {
		return
	}
	customer = entity.Customer{
		Email:        request.Email,
		Name:         request.Name,
		Password:     request.Password,
		RegisteredAt: request.RegisteredAt,
	}
	return customer, nil
}

func (repository *CustomerRepositoryImpl) LoginCustomer(ctx context.Context, email string) (id uuid.UUID, passwordHashed string, err error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	query := "SELECT id, password FROM customers WHERE email = $1"
	err = tx.QueryRowContext(ctx, query, email).Scan(&id, &passwordHashed)
	if err != nil {
		if err != nil {
			return
		}
	}

	return
}

func (repository *CustomerRepositoryImpl) UpdateCustomer(ctx context.Context, request entity.Customer) (customer entity.Customer, err error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	query := "UPDATE customers SET name=$1, updated_at=$2 WHERE email=$3"
	_, err = tx.ExecContext(ctx, query, request.Name, request.UpdatedAt, request.Email)
	if err != nil {
		return
	}
	queryResult := "SELECT name, registered_at, updated_at FROM customers WHERE email=$1"
	err = tx.QueryRowContext(ctx, queryResult, request.Email).Scan(&customer.Name, &customer.RegisteredAt, &customer.UpdatedAt)
	if err != nil {
		return
	}

	customer.Email = request.Email

	return
}

func (repository *CustomerRepositoryImpl) DeleteCustomer(ctx context.Context, email string, deleteTime int32) (err error) {
	tx, err := repository.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	query := "UPDATE customers SET deleted_at=$1 WHERE email=$2"
	_, err = tx.ExecContext(ctx, query, deleteTime, email)
	if err != nil {
		return
	}
	return
}
