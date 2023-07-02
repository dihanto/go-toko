package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dihanto/go-toko/model/entity"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepositoryImpl() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (repository *OrderRepositoryImpl) AddOrder(ctx context.Context, tx *sql.Tx, request entity.Order) (order entity.Order, err error) {
	queryOrder := "INSERT INTO orders (id_product, id_customer, quantity) VALUES ($1, $2, $3) RETURNING id"
	err = tx.QueryRowContext(ctx, queryOrder, request.IdProduct, request.IdCustomer, request.Quantity).Scan(&request.Id)
	if err != nil {
		return
	}

	var price int
	var resultQuantity string
	queryProduct := "UPDATE products SET quantity = CASE WHEN (quantity - $1) < 0 THEN quantity ELSE quantity - $1 END WHERE id = $2 RETURNING CASE WHEN (quantity - $1) < 0 THEN 'Quantity cannot be less then 0' ELSE 'Success' END AS result, price"
	tx.QueryRowContext(ctx, queryProduct, request.Quantity, request.IdProduct).Scan(&resultQuantity, &price)
	if resultQuantity != "Success" {
		return order, errors.New(resultQuantity)
	}

	var resultBalance string
	totalPrice := price * request.Quantity
	queryWallet := "Update wallet SET balance= CASE WHEN (balance-$1) < 0 THEN balance ELSE balance - $1 END WHERE id_customer=$2 RETURNING CASE WHEN (balance - $1) < 0 THEN 'Balance cannot be less than 0' ELSE 'Success' END AS result"
	tx.QueryRowContext(ctx, queryWallet, totalPrice, request.IdCustomer).Scan(&resultBalance)
	if resultBalance != "Success" {
		return order, errors.New(resultBalance)
	}

	order = entity.Order{
		Id:         request.Id,
		IdProduct:  request.IdProduct,
		IdCustomer: request.IdCustomer,
		Quantity:   request.Quantity,
	}

	return

}

func (repository *OrderRepositoryImpl) FindOrder(ctx context.Context, tx *sql.Tx, id int) (order entity.Order, product entity.Product, customerName string, err error) {
	query := "SELECT o.id_product, o.id_customer, o.quantity, p.name, p.price, c.name FROM orders o JOIN products p ON o.id_product = p.id JOIN customers c ON o.id_customer = c.id WHERE o.id=$1"
	err = tx.QueryRowContext(ctx, query, id).Scan(&order.IdProduct, &order.IdCustomer, &order.Quantity, &product.Name, &product.Price, &customerName)
	if err != nil {
		return
	}

	order.Id = id

	return
}
