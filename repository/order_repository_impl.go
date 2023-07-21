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

func (repository *OrderRepositoryImpl) AddOrder(ctx context.Context, tx *sql.Tx, orderRequest entity.Order, orderDetailRequest entity.OrderDetail) (order entity.Order, orderDetail entity.OrderDetail, err error) {
	queryOrder := "INSERT INTO orders (id_customer, ordered_at) VALUES ($1, $2) RETURNING id"
	err = tx.QueryRowContext(ctx, queryOrder, orderRequest.IdCustomer, orderRequest.OrderedAt).Scan(&orderRequest.Id)
	if err != nil {
		return
	}

	queryOrderDetail := "INSERT INTO order_details (id_product, quantity, id_order) VALUES ($1, $2, $3)"
	_, err = tx.ExecContext(ctx, queryOrderDetail, orderDetailRequest.IdProduct, orderDetailRequest.Quantity, orderRequest.Id)
	if err != nil {
		return
	}

	var price int
	var resultQuantity string
	queryProduct := `UPDATE products SET quantity = CASE WHEN (quantity - $1) < 0 THEN quantity ELSE quantity - $1 END, updated_at=$2 WHERE id = $3
	RETURNING CASE WHEN (quantity - $1) < 0 THEN 'Quantity cannot be less then 0' ELSE 'Success' END AS result, price`
	err = tx.QueryRowContext(ctx, queryProduct, orderDetailRequest.Quantity, orderRequest.OrderedAt, orderDetailRequest.IdProduct).Scan(&resultQuantity, &price)
	if err != nil {
		return
	}
	if resultQuantity != "Success" {
		return order, orderDetail, errors.New(resultQuantity)
	}

	var resultBalance string
	totalPrice := price * orderDetailRequest.Quantity
	queryWallet := "Update wallets SET balance = CASE WHEN (balance-$1) < 0 THEN balance ELSE balance - $1 END, updated_at=$2 WHERE id_customer=$3 RETURNING CASE WHEN (balance - $1) < 0 THEN 'Balance cannot be less than 0' ELSE 'Success' END AS result"
	err = tx.QueryRowContext(ctx, queryWallet, totalPrice, orderRequest.OrderedAt, orderRequest.IdCustomer).Scan(&resultBalance)
	if err != nil {
		return
	}
	if resultBalance != "Success" {
		return order, orderDetail, errors.New(resultBalance)
	}

	order = entity.Order{
		Id:         orderRequest.Id,
		IdCustomer: orderRequest.IdCustomer,
		OrderedAt:  orderRequest.OrderedAt,
	}

	orderDetail = entity.OrderDetail{
		IdOrder:   orderRequest.Id,
		IdProduct: orderDetailRequest.IdProduct,
		Quantity:  orderDetailRequest.Quantity,
	}

	return order, orderDetail, nil

}

func (repository *OrderRepositoryImpl) FindOrder(ctx context.Context, tx *sql.Tx, id int) (order entity.Order, orderDetail entity.OrderDetail, product entity.Product, customerName string, err error) {
	queryOrder := "SELECT o.id_customer,  o.ordered_at, c.name FROM orders o JOIN customers c ON o.id_customer = c.id WHERE o.id=$1"
	err = tx.QueryRowContext(ctx, queryOrder, id).Scan(&order.IdCustomer, &order.OrderedAt, &customerName)
	if err != nil {
		return
	}
	order.Id = id

	queryOrderDetail := "SELECT od.id_product, od.quantity, p.name, p.price FROM order_details od JOIN products p ON od.id_product = p.id WHERE od.id_order=$1"
	err = tx.QueryRowContext(ctx, queryOrderDetail, id).Scan(&orderDetail.IdProduct, &orderDetail.Quantity, &product.Name, &product.Price)
	if err != nil {
		return
	}

	return
}
