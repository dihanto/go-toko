package repository

import (
	"context"
	"log"
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

	mock.ExpectBegin()
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

func TestSellerLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}

	emailPayload := "luffy@onepiece.com"
	id := uuid.NewString()

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id, password FROM sellers WHERE email=\\$1").WithArgs(emailPayload).
		WillReturnRows(sqlmock.NewRows([]string{"id", "password"}).AddRow(id, "kjsdflkjskdf"))

	repo := NewSellerRepositoryImpl(db)

	_, _, err = repo.LoginSeller(context.Background(), emailPayload)
	assert.NoError(t, err)
}

func TestUpdateSeller(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}

	sellerPayload := entity.Seller{
		Name:      "Luffy",
		Email:     "luffy@onepiece.com",
		UpdatedAt: int32(time.Now().Unix()),
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE sellers SET name=\\$1, updated_at=\\$2 WHERE email=\\$3").
		WithArgs(sellerPayload.Name, sellerPayload.UpdatedAt, sellerPayload.Email).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT name, registered_at, updated_at FROM sellers WHERE email=\\$1").
		WithArgs(sellerPayload.Email).WillReturnRows(sqlmock.NewRows([]string{"name", "registered_at", "updated_at"}).AddRow("Luffy", "129839122", "47213019"))

	repo := NewSellerRepositoryImpl(db)
	_, err = repo.UpdateSeller(context.Background(), sellerPayload)
	assert.NoError(t, err)
}

func TestDeleteSeller(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}

	deleteTime := int32(time.Now().Unix())
	email := "luffy@onepiece.com"

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE sellers SET deleted_at=\\$1 WHERE email=\\$2").WithArgs(deleteTime, email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewSellerRepositoryImpl(db)
	err = repo.DeleteSeller(context.Background(), deleteTime, email)
	assert.NoError(t, err)
}
