package repository

import (
	"context"
	"database/sql"

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

	queryProduct := "UPDATE products SET quantity=quantity-$1 WHERE id=$2"
	_, err = tx.ExecContext(ctx, queryProduct, request.Quantity, request.IdProduct)
	if err != nil {
		return
	}
	var price int
	queryPrice := "SELECT price FROM products WHERE id=$1"
	err = tx.QueryRowContext(ctx, queryPrice, request.IdProduct).Scan(&price)
	if err != nil {
		return
	}
	totalPrice := price * request.Quantity

	queryWallet := "Update wallet SET balance=balance-$1 WHERE id_customer=$2"
	_, err = tx.ExecContext(ctx, queryWallet, totalPrice, request.IdCustomer)
	if err != nil {
		return
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
