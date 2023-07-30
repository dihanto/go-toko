package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-toko/model/entity"
	"github.com/google/uuid"
)

type SellerRepositoryImpl struct {
	Database *sql.DB
}

func NewSellerRepositoryImpl(database *sql.DB) SellerRepository {
	return &SellerRepositoryImpl{
		Database: database,
	}
}

func (repository *SellerRepositoryImpl) RegisterSeller(ctx context.Context, request entity.Seller) (seller entity.Seller, err error) {
	query := `INSERT INTO sellers (id, email, name, password, registered_at) VALUES ($1, $2, $3, $4, $5)`
	_, err = repository.Database.ExecContext(ctx, query, request.Id, request.Email, request.Name, request.Password, request.RegisteredAt)
	if err != nil {
		return
	}
	seller = entity.Seller{
		Email:        request.Email,
		Name:         request.Name,
		Password:     request.Password,
		RegisteredAt: request.RegisteredAt,
	}
	return
}

func (repository *SellerRepositoryImpl) LoginSeller(ctx context.Context, email string) (id uuid.UUID, password string, err error) {
	query := "SELECT id, password FROM sellers WHERE email=$1"
	err = repository.Database.QueryRowContext(ctx, query, email).Scan(&id, &password)
	if err != nil {
		return
	}

	return
}

func (repository *SellerRepositoryImpl) UpdateSeller(ctx context.Context, request entity.Seller) (seller entity.Seller, err error) {
	query := "UPDATE sellers SET name=$1, updated_at=$2 WHERE email=$3"
	_, err = repository.Database.ExecContext(ctx, query, request.Name, request.UpdatedAt, request.Email)
	if err != nil {
		return
	}
	queryResult := "SELECT name, registered_at, updated_at FROM sellers WHERE email=$1"
	rows, err := repository.Database.QueryContext(ctx, queryResult, request.Email)
	if err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&seller.Name, &seller.RegisteredAt, &seller.UpdatedAt)
		if err != nil {
			return
		}
	}

	seller.Email = request.Email

	return
}

func (repository *SellerRepositoryImpl) DeleteSeller(ctx context.Context, deleteTime int32, email string) (err error) {
	query := "UPDATE sellers SET deleted_at=$1 WHERE email=$2"
	_, err = repository.Database.ExecContext(ctx, query, deleteTime, email)
	if err != nil {
		return
	}
	return
}
