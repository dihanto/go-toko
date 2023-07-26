package usecase

import (
	"context"
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
	Validate   *validator.Validate
	Timeout    time.Duration
}

func NewOrderUsecaseImpl(repository repository.OrderRepository, validate *validator.Validate, timeout time.Duration) OrderUsecase {
	return &OrderUsecaseImpl{
		Repository: repository,
		Validate:   validate,
		Timeout:    timeout,
	}
}

func (usecase *OrderUsecaseImpl) AddOrder(ctx context.Context, request request.AddOrder) (order response.AddOrder, err error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.Timeout*time.Second)
	defer cancel()

	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}

	requestOrder := entity.Order{
		IdCustomer: request.IdCustomer,
		OrderedAt:  int32(time.Now().Unix()),
	}

	requestOrderDetail := entity.OrderDetail{
		IdProduct: request.IdProduct,
		Quantity:  request.Quantity,
	}

	responseOrder, responseOrderDetail, err := usecase.Repository.AddOrder(ctx, requestOrder, requestOrderDetail)
	if err != nil {
		return
	}

	order = helper.ToResponseAddOrder(responseOrder, responseOrderDetail)

	return
}

func (usecase *OrderUsecaseImpl) FindOrder(ctx context.Context, id int) (orderDetail response.FindOrder, err error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.Timeout*time.Second)
	defer cancel()

	err = usecase.Validate.Var(id, "required")
	if err != nil {
		return
	}

	responseOrder, responseOrderDetail, product, customerName, err := usecase.Repository.FindOrder(ctx, id)
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
		Id:         responseOrder.Id,
		IdProduct:  responseOrderDetail.IdProduct,
		IdCustomer: responseOrder.IdCustomer,
		Quantity:   responseOrderDetail.Quantity,
		OrderedAt:  time.Unix(int64(responseOrder.OrderedAt), 0),
		TotalPrice: responseOrderDetail.Quantity * product.Price,
		Product:    productDetail,
		Customer:   customerDetail,
	}

	return
}
