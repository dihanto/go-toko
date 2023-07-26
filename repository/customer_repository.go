package repository

import (
	"context"

	"github.com/dihanto/go-toko/model/entity"
	"github.com/google/uuid"
)

type CustomerRepository interface {
	RegisterCustomer(ctx context.Context, request entity.Customer) (customer entity.Customer, err error)
	LoginCustomer(ctx context.Context, email string) (id uuid.UUID, passwordHashed string, err error)
	UpdateCustomer(ctx context.Context, request entity.Customer) (customer entity.Customer, err error)
	DeleteCustomer(ctx context.Context, email string, deleteTime int32) (err error)
}
