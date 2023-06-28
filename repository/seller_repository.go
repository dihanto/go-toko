package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-toko/model/entity"
	"github.com/google/uuid"
)

type SellerRepository interface {
	RegisterSeller(ctx context.Context, tx *sql.Tx, request entity.Seller) (seller entity.Seller, err error)
	LoginSeller(ctx context.Context, tx *sql.Tx, email string) (id uuid.UUID, password string, err error)
	UpdateSeller(ctx context.Context, tx *sql.Tx, request entity.Seller) (seller entity.Seller, err error)
	DeleteSeller(ctx context.Context, tx *sql.Tx, deleteTime int32, email string) (err error)
}
