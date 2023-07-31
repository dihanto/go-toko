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

func TestAddProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}

	productPayload := entity.Product{
		IdSeller:  uuid.New(),
		Name:      "Luffy",
		Price:     1000,
		Quantity:  20,
		CreatedAt: int32(time.Now().Unix()),
	}

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO products \\(id_seller, name, price, quantity, created_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\) RETURNING id").
		WithArgs(productPayload.IdSeller, productPayload.Name, productPayload.Price, productPayload.Quantity, productPayload.CreatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	repo := NewProductRepositoryImpl(db)
	_, err = repo.AddProduct(context.Background(), productPayload)
	assert.NoError(t, err)
}

func TestGetProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id, name, price FROM products WHERE deleted_at IS NULL").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).AddRow(1, "jeruk", 1000))

	repo := NewProductRepositoryImpl(db)
	_, err = repo.GetProduct(context.Background())
	assert.NoError(t, err)
}

func TestFindProductById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}

	id := 1
	idSeller := uuid.NewString()

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id_seller, name, price, quantity, created_at, updated_at FROM products WHERE id=\\$1").
		WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"id_seller", "name", "price", "quantity", "created_at", "updated_at"}).
		AddRow(idSeller, "jeruk", 1000, 10, "12381237", "12345678"))

	repo := NewProductRepositoryImpl(db)
	_, err = repo.FindById(context.Background(), id)
	assert.NoError(t, err)
}

func TestUpdateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}

	productPayload := entity.Product{
		Id:        1,
		Name:      "jeruk",
		Price:     1000,
		Quantity:  10,
		UpdatedAt: int32(time.Now().Unix()),
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products SET name=\\$1, price=\\$2, quantity=\\$3, updated_at=\\$4 WHERE id=\\$5").
		WithArgs(productPayload.Name, productPayload.Price, productPayload.Quantity, productPayload.UpdatedAt, productPayload.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT name, price, quantity, created_at, updated_at FROM products WHERE id=\\$1").WithArgs(productPayload.Id).
		WillReturnRows(sqlmock.NewRows([]string{"name", "price", "quantity", "created_at", "updated_at"}).AddRow("Luffy", 1000, 10, "12345678", "12345678"))

	repo := NewProductRepositoryImpl(db)
	_, err = repo.UpdateProduct(context.Background(), productPayload)
	assert.NoError(t, err)
}

func TestDeleteProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}

	id := 1
	deleteTime := int32(time.Now().Unix())

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products SET deleted_at=\\$1 WHERE id=\\$2").WithArgs(deleteTime, id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewProductRepositoryImpl(db)
	err = repo.DeleteProduct(context.Background(), deleteTime, id)
	assert.NoError(t, err)
}

func TestSearchProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}

	search := "jeruk"
	offset := 5
	limit := 10
	query := `SELECT id, name, price, quantity FROM products
          WHERE name LIKE '%' || $1 || '%'
          OR CAST(price AS TEXT) LIKE '%' || $1 || '%'
          OR CAST(quantity AS TEXT) LIKE '%' || $1 || '%'
          ORDER BY created_at DESC
          LIMIT $2 OFFSET $3`
	queryCount := `SELECT COUNT(*) FROM products
		  WHERE name LIKE '%' || $1 || '%'
		  OR CAST(price AS TEXT) LIKE '%' || $1 || '%'
		  OR CAST(quantity AS TEXT) LIKE '%' || $1 || '%'`

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(search, limit, offset).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "LIMIT", "OFFSET"}).AddRow(1, "jeruk", 10, 5))

	mock.ExpectQuery(regexp.QuoteMeta(queryCount)).
		WithArgs(search).WillReturnRows(sqlmock.NewRows([]string{"COUNT"}).AddRow(1))

	repo := NewProductRepositoryImpl(db)
	_, _, err = repo.Search(context.Background(), search, offset, limit)
	assert.NoError(t, err)

}

func TestAddProductToWishlist(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}

	id := 1
	idCustomer := uuid.New()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE products SET wishlist=wishlist+1 WHERE id=$1")).WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO wishlist_details \\(product_id, customer_id\\) VALUES \\(\\$1, \\$2\\)").
		WithArgs(id, idCustomer).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT name, price, quantity, wishlist FROM products WHERE id=\\$1").
		WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"name", "price", "quantity", "wishlist"}).AddRow("luffy", 1000, 10, 2))

	repo := NewProductRepositoryImpl(db)

	_, err = repo.AddProductToWishlist(context.Background(), id, idCustomer)
	assert.NoError(t, err)
}

func TestDeleteProductFromWishlist(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}

	id := 1
	idCustomer := uuid.New()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE products SET wishlist=wishlist-1 WHERE id=$1")).WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("DELETE FROM wishlist_details WHERE customer_id=\\$1").WithArgs(idCustomer).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT name, price, quantity, wishlist FROM products WHERE id=\\$1").WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"name", "price", "quantity", "wishlist"}).AddRow("jeruk", 1000, 10, 1))

	repo := NewProductRepositoryImpl(db)
	_, err = repo.DeleteProductFromWishlist(context.Background(), id, idCustomer)
	assert.NoError(t, err)
}
