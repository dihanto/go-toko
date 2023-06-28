package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-toko/model/entity"
	"github.com/google/uuid"
)

type CustomerRepository interface {
	RegisterCustomer(ctx context.Context, tx *sql.Tx, request entity.Customer) (customer entity.Customer, err error)
	LoginCustomer(ctx context.Context, tx *sql.Tx, email string) (id uuid.UUID, passwordHashed string, err error)
	UpdateCustomer(ctx context.Context, tx *sql.Tx, request entity.Customer) (customer entity.Customer, err error)
	DeleteCustomer(ctx context.Context, tx *sql.Tx, email string, deleteTime int32) (err error)
}
