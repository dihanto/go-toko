package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/dihanto/go-toko/helper"
	"github.com/dihanto/go-toko/model/entity"
	"github.com/dihanto/go-toko/model/web/request"
	"github.com/dihanto/go-toko/model/web/response"
	"github.com/dihanto/go-toko/repository"
	"github.com/go-playground/validator/v10"
)

type OrderUsecaseImpl struct {
	Repository repository.OrderRepository
	Db         *sql.DB
	Validate   *validator.Validate
	Timeout    int
}

func NewOrderUsecaseImpl(repository repository.OrderRepository, db *sql.DB, validate *validator.Validate, timeout int) OrderUsecase {
	return &OrderUsecaseImpl{
		Repository: repository,
		Db:         db,
		Validate:   validate,
		Timeout:    timeout,
	}
}

func (usecase *OrderUsecaseImpl) AddOrder(ctx context.Context, request request.AddOrder) (order response.AddOrder, err error) {
	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}

	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	requestRepo := entity.Order{
		IdProduct:  request.IdProduct,
		IdCustomer: request.IdCustomer,
		Quantity:   request.Quantity,
		OrderedAt:  int32(time.Now().Unix()),
	}

	response, err := usecase.Repository.AddOrder(ctx, tx, requestRepo)
	if err != nil {
		return
	}

	order = helper.ToResponseAddOrder(response)

	return
}

func (usecase *OrderUsecaseImpl) FindOrder(ctx context.Context, id int) (orderDetail response.FindOrder, err error) {
	err = usecase.Validate.Var(id, "required")
	if err != nil {
		return
	}

	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	order, product, customerName, err := usecase.Repository.FindOrder(ctx, tx, id)
	if err != nil {
		return
	}

	productDetail := response.ProductOrder{
		Name:  product.Name,
		Price: product.Price,
	}
	customerDetail := response.CustomerOrder{
		Name: customerName,
	}
	orderDetail = response.FindOrder{
		Id:         order.Id,
		IdProduct:  order.IdProduct,
		IdCustomer: order.IdCustomer,
		Quantity:   order.Quantity,
		OrderedAt:  time.Unix(int64(order.OrderedAt), 0),
		TotalPrice: order.Quantity * product.Price,
		Product:    productDetail,
		Customer:   customerDetail,
	}

	return
}
