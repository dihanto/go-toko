package usecase

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/dihanto/go-toko/helper"
	"github.com/dihanto/go-toko/model/entity"
	"github.com/dihanto/go-toko/model/web/request"
	"github.com/dihanto/go-toko/model/web/response"
	"github.com/dihanto/go-toko/repository"
	"github.com/go-playground/validator/v10"
)

type ProductUsecaseImpl struct {
	Repository repository.ProductRepository
	Db         *sql.DB
	Validate   *validator.Validate
	Timeout    time.Duration
}

func NewProductUsecaseImpl(repository repository.ProductRepository, db *sql.DB, validate *validator.Validate, timeout time.Duration) ProductUsecase {
	return &ProductUsecaseImpl{
		Repository: repository,
		Db:         db,
		Validate:   validate,
		Timeout:    timeout,
	}
}

func (usecase *ProductUsecaseImpl) AddProduct(ctx context.Context, request request.AddProduct) (product response.AddProduct, err error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.Timeout*time.Second)
	defer cancel()

	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}

	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	requestRepo := entity.Product{
		IdSeller:  request.IdSeller,
		Name:      request.Name,
		Price:     request.Price,
		Quantity:  request.Quantity,
		CreatedAt: int32(time.Now().Unix()),
	}
	response, err := usecase.Repository.AddProduct(ctx, tx, requestRepo)
	if err != nil {
		return
	}

	product = helper.ToResponseAddProduct(response)

	return
}

func (usecase *ProductUsecaseImpl) GetProduct(ctx context.Context) (products []response.GetProduct, err error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.Timeout*time.Second)
	defer cancel()

	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	responses, err := usecase.Repository.GetProduct(ctx, tx)
	if err != nil {
		return
	}

	for _, product := range responses {
		response := response.GetProduct{
			Id:    product.Id,
			Name:  product.Name,
			Price: product.Price,
		}
		products = append(products, response)
	}
	return
}

func (usecase *ProductUsecaseImpl) FindById(ctx context.Context, id int) (product response.FindById, err error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.Timeout*time.Second)
	defer cancel()

	err = usecase.Validate.Var(id, "required")
	if err != nil {
		return
	}

	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	response, err := usecase.Repository.FindById(ctx, tx, id)
	if err != nil {
		return
	}
	product = helper.ToResponseFindById(response)
	return
}
func (usecase *ProductUsecaseImpl) UpdateProduct(ctx context.Context, request request.UpdateProduct) (product response.UpdateProduct, err error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.Timeout*time.Second)
	defer cancel()

	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}

	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	requestRepo := entity.Product{
		Id:        request.Id,
		Name:      request.Name,
		Price:     request.Price,
		Quantity:  request.Quantity,
		UpdatedAt: int32(time.Now().Unix()),
	}

	response, err := usecase.Repository.UpdateProduct(ctx, tx, requestRepo)
	if err != nil {
		return
	}

	product = helper.ToResponseUpdateProduct(response)

	return
}

func (usecase *ProductUsecaseImpl) DeleteProduct(ctx context.Context, id int) (err error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.Timeout*time.Second)
	defer cancel()

	err = usecase.Validate.Var(id, "required")
	if err != nil {
		return
	}

	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	deleteTime := int32(time.Now().Unix())
	err = usecase.Repository.DeleteProduct(ctx, tx, deleteTime, id)
	if err != nil {
		return
	}
	return
}

func (usecase *ProductUsecaseImpl) FindByName(ctx context.Context, search string, offset int, limit int) (productsWithPagination response.FindByName, err error) {
	ctx, cancel := context.WithTimeout(ctx, 900*time.Second)
	defer cancel()

	err = usecase.Validate.Var(search, "required")
	if err != nil {
		return
	}

	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	responses, countString, err := usecase.Repository.FindByName(ctx, tx, search, offset, limit)
	if err != nil {
		return
	}

	var products []response.FindByNameProduct
	for _, product := range responses {
		response := response.FindByNameProduct{
			Id:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		}
		products = append(products, response)
	}

	count, err := strconv.Atoi(countString)
	if err != nil {
		return
	}

	currentPage := ((count / limit) + 1) - ((count - offset + 1) / limit)
	lastPage := (count / limit) + 1
	pagination := response.Pagination{
		Total:        count,
		PerPage:      limit,
		CurrenntPage: currentPage,
		LastPage:     lastPage,
	}

	productsWithPagination = response.FindByName{
		Product:    products,
		Pagination: pagination,
	}

	return productsWithPagination, nil
}

func (usecase *ProductUsecaseImpl) AddProductToWishlist(ctx context.Context, request request.AddProductToWishlist) (product response.AddProductToWishlist, err error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.Timeout*time.Second)
	defer cancel()

	err = usecase.Validate.Struct(request)
	if err != nil {
		return
	}
	productId := strconv.Itoa(request.ProductId)
	err = usecase.Validate.Var(request.CustomerId, "wishlist="+productId)
	if err != nil {
		return
	}

	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx, &err)

	response, err := usecase.Repository.AddProductToWishlist(ctx, tx, request.ProductId, request.CustomerId)
	if err != nil {
		return
	}

	product = helper.ToReponseAddProductToWishlist(response)

	return product, nil
}

func (usecase *ProductUsecaseImpl) DeleteProductFromWishlist(ctx context.Context, request request.DeleteProductFromWishlist) (product response.DeleteProductFromWishlist, err error) {
	panic("not implemented") // TODO: Implement
}
