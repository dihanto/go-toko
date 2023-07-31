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

	mock.ExpectBegin()
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

func TestFindOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}

	id := 1
	idCustomer := uuid.NewString()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT o.id_customer,  o.ordered_at, c.name FROM orders o JOIN customers c ON o.id_customer = c.id WHERE o.id=$1")).
		WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"id_customer", "quantity", "name"}).AddRow(idCustomer, 2, "jeruk"))

	mock.ExpectQuery(regexp.QuoteMeta("SELECT od.id_product, od.quantity, p.name, p.price FROM order_details od JOIN products p ON od.id_product = p.id WHERE od.id_order=$1")).
		WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"id_product", "quantity", "name", "price"}).AddRow("1", 2, "jeruk", 1000))

	repo := NewOrderRepositoryImpl(db)
	_, _, _, _, err = repo.FindOrder(context.Background(), id)
	assert.NoError(t, err)
}
