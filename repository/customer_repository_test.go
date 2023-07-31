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

func TestRegisterCustomer(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Println(err)
	}

	id := uuid.New()
	customerPayload := entity.Customer{
		Id:           id,
		Name:         "Luffy",
		Email:        "luffy@onepiece.com",
		Password:     "skdjfsjfk",
		RegisteredAt: int32(time.Now().Unix()),
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO customers \\(id, email, name, password, registered_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\);").
		WithArgs(customerPayload.Id, customerPayload.Email, customerPayload.Name, customerPayload.Password, customerPayload.RegisteredAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewCustomerRepositoryImpl(db)
	_, err = repo.RegisterCustomer(context.Background(), customerPayload)
	assert.NoError(t, err)

}

func TestLoginCustomer(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Println(err)
	}
	email := "luffy@onepiece.com"
	id := uuid.NewString()

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id, password FROM customers WHERE email = \\$1").
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "password"}).
			AddRow(id, "ksadjfkasjdfkj"))

	repo := NewCustomerRepositoryImpl(db)
	_, _, err = repo.LoginCustomer(context.Background(), email)
	assert.NoError(t, err)

}

func TestUpdateCustomer(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Println(err)
	}

	customerPayload := entity.Customer{
		Name:      "Luffy",
		Email:     "luffy@onepiece.com",
		UpdatedAt: int32(time.Now().Unix()),
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE customers SET name=\\$1, updated_at=\\$2 WHERE email=\\$3").
		WithArgs(customerPayload.Name, customerPayload.UpdatedAt, customerPayload.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT name, registered_at, updated_at FROM customers WHERE email=\\$1").
		WithArgs(customerPayload.Email).
		WillReturnRows(sqlmock.NewRows([]string{"name", "registered_at", "updated_at"}).
			AddRow("luffy", "12831780", "12303091"))

	repo := NewCustomerRepositoryImpl(db)

	_, err = repo.UpdateCustomer(context.Background(), customerPayload)
	assert.NoError(t, err)
}

func TestDeleteCustomer(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Println(err)
	}

	deleteTime := int32(time.Now().Unix())
	email := "luffy@onepiece.com"

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE customers SET deleted_at=\\$1 WHERE email=\\$2").
		WithArgs(deleteTime, email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewCustomerRepositoryImpl(db)
	err = repo.DeleteCustomer(context.Background(), email, deleteTime)
	assert.NoError(t, err)

}
