package repository

import (
	"context"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/dihanto/go-toko/model/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestAddOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}

	orderPayload := entity.Order{
		Id:         1,
		IdCustomer: uuid.New(),
		OrderedAt:  int32(time.Now().Unix()),
	}

	orderDetailPayload := entity.OrderDetail{
		IdProduct: 1,
		Quantity:  2,
		Id:        1,
	}

	totalPricePayload := 4000

	mock.ExpectQuery("INSERT INTO orders \\(id_customer, ordered_at\\) VALUES \\(\\$1, \\$2\\) RETURNING id").
		WithArgs(orderPayload.IdCustomer, orderPayload.OrderedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectExec("INSERT INTO order_details \\(id_product, quantity, id_order\\) VALUES \\(\\$1, \\$2, \\$3\\)").
		WithArgs(orderDetailPayload.IdProduct, orderDetailPayload.Quantity, orderPayload.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery(regexp.QuoteMeta("UPDATE products SET quantity = CASE WHEN (quantity - $1) < 0 THEN quantity ELSE quantity - $1 END, updated_at=$2 WHERE id = $3 RETURNING CASE WHEN (quantity - $1) < 0 THEN 'Quantity cannot be less then 0' ELSE 'Success' END AS result, price")).
		WithArgs(orderDetailPayload.Quantity, orderPayload.OrderedAt, orderDetailPayload.IdProduct).
		WillReturnRows(sqlmock.NewRows([]string{"result", "price"}).AddRow("Success", 2000))

	mock.ExpectQuery(regexp.QuoteMeta("UPDATE wallets SET balance = CASE WHEN (balance-$1) < 0 THEN balance ELSE balance - $1 END, updated_at=$2 WHERE id_customer=$3 RETURNING CASE WHEN (balance - $1) < 0 THEN 'Balance cannot be less than 0' ELSE 'Success' END AS result")).
		WithArgs(totalPricePayload, orderPayload.OrderedAt, orderPayload.IdCustomer).
		WillReturnRows(sqlmock.NewRows([]string{"result"}).AddRow("Success"))

	repo := NewOrderRepositoryImpl(db)

	_, _, err = repo.AddOrder(context.Background(), orderPayload, orderDetailPayload)
	assert.NoError(t, err)
}
