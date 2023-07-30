package repository

import (
	"context"
	"testing"
	"time"

	"github.com/dihanto/go-toko/model/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSellerRegister(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock: %s", err)
	}
	defer db.Close()
	seller := entity.Seller{
		Id:           uuid.New(),
		Email:        "luffy@onepiece.com",
		Name:         "Luffy",
		Password:     "lskjdflksjdf",
		RegisteredAt: int32(time.Now().Unix()),
	}

	mock.ExpectExec("INSERT INTO sellers \\(id, email, name, password, registered_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\)").
		WithArgs(seller.Id, seller.Email, seller.Name, seller.Password, seller.RegisteredAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewSellerRepositoryImpl(db)

	_, err = repo.RegisterSeller(context.TODO(), seller)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
