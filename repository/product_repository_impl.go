package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-toko/model/entity"
	"github.com/google/uuid"
)

type ProductRepositoryImpl struct {
}

func NewProductRepositoryImpl() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) AddProduct(ctx context.Context, tx *sql.Tx, request entity.Product) (product entity.Product, err error) {
	query := "INSERT INTO products (id_seller, name, price, quantity, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id"
	err = tx.QueryRowContext(ctx, query, request.IdSeller, request.Name, request.Price, request.Quantity, request.CreatedAt).Scan(&request.Id)
	if err != nil {
		return
	}

	product = entity.Product{
		Id:        request.Id,
		IdSeller:  request.IdSeller,
		Name:      request.Name,
		Price:     request.Price,
		CreatedAt: request.CreatedAt,
		Quantity:  request.Quantity,
	}
	return
}

func (repository *ProductRepositoryImpl) GetProduct(ctx context.Context, tx *sql.Tx) (products []entity.Product, err error) {
	query := "SELECT id, name, price FROM products WHERE deleted_at IS NULL"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Price)
		if err != nil {
			return
		}
		products = append(products, product)
	}
	return
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (product entity.Product, err error) {
	query := "Select id_seller, name, price, quantity, created_at, updated_at FROM products WHERE id=$1"
	err = tx.QueryRowContext(ctx, query, id).Scan(&product.IdSeller, &product.Name, &product.Price, &product.Quantity, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return
	}

	product.Id = id

	return
}

func (repository *ProductRepositoryImpl) UpdateProduct(ctx context.Context, tx *sql.Tx, request entity.Product) (product entity.Product, err error) {
	query := "UPDATE products SET name=$1, price=$2, quantity=$3, updated_at=$4 WHERE id=$5"
	_, err = tx.ExecContext(ctx, query, request.Name, request.Price, request.Quantity, request.UpdatedAt, request.Id)
	if err != nil {
		return
	}
	queryProduct := "SELECT name, price, quantity, created_at, updated_at FROM products WHERE id=$1"
	err = tx.QueryRowContext(ctx, queryProduct, request.Id).Scan(&product.Name, &product.Price, &product.Quantity, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return
	}
	product.Id = request.Id

	return
}

func (repository *ProductRepositoryImpl) DeleteProduct(ctx context.Context, tx *sql.Tx, deleteTime int32, id int) (err error) {
	query := "UPDATE products SET deleted_at=$1 WHERE id=$2"
	_, err = tx.ExecContext(ctx, query, deleteTime, id)
	if err != nil {
		return
	}
	return
}

func (repository *ProductRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, search string, offset int, limit int) (products []entity.Product, count string, err error) {

	query := `SELECT id, name, price, quantity FROM products
			  WHERE name LIKE $1
			  OR CAST(price AS TEXT) LIKE $1
			  OR CAST(quantity AS TEXT) LIKE $1
			  ORDER BY created_at DESC
			  LIMIT $2 OFFSET $3`

	rows, err := tx.QueryContext(ctx, query, "%"+search+"%", limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity)
		if err != nil {
			return
		}
		products = append(products, product)
	}

	queryCount := `SELECT COUNT(*) FROM products
					WHERE name LIKE $1
					OR CAST(price AS TEXT) LIKE $1
					OR CAST(quantity AS TEXT) LIKE $1`
	rowsCount, err := tx.QueryContext(ctx, queryCount, "%"+search+"%")
	if err != nil {
		return
	}
	defer rowsCount.Close()

	if rowsCount.Next() {
		err = rowsCount.Scan(&count)
		if err != nil {
			return
		}
	}

	return products, count, nil
}

func (repository *ProductRepositoryImpl) AddProductToWishlist(ctx context.Context, tx *sql.Tx, productId int, customerId uuid.UUID) (product entity.Product, err error) {
	query := "UPDATE products SET wishlist=wishlist+1 WHERE id=$1"
	_, err = tx.ExecContext(ctx, query, productId)
	if err != nil {
		return
	}

	queryWishlist := "INSERT INTO wishlist_details (product_id, customer_id) VALUES ($1, $2)"
	_, err = tx.ExecContext(ctx, queryWishlist, productId, customerId)
	if err != nil {
		return
	}

	queryProduct := "SELECT name, price, quantity, wishlist FROM products WHERE id=$1"
	err = tx.QueryRowContext(ctx, queryProduct, productId).Scan(&product.Name, &product.Price, &product.Quantity, &product.Wishlist)
	if err != nil {
		return
	}
	product.Id = productId

	return product, nil
}

func (repository *ProductRepositoryImpl) DeleteProductFromWishlist(ctx context.Context, tx *sql.Tx, productId int, customerId uuid.UUID) (product entity.Product, err error) {
	query := "UPDATE products SET wishlist=wishlist-1 WHERE id=$1"
	_, err = tx.ExecContext(ctx, query, productId)
	if err != nil {
		return
	}

	queryWishlist := "DELETE FROM wishlist_details WHERE customer_id=$1"
	_, err = tx.ExecContext(ctx, queryWishlist, customerId)
	if err != nil {
		return
	}

	queryProduct := "SELECT name, price, quantity, wishlist FROM products WHERE id=$1"
	err = tx.QueryRowContext(ctx, queryProduct, productId).Scan(&product.Name, &product.Price, &product.Quantity, &product.Wishlist)
	if err != nil {
		return
	}
	product.Id = productId

	return product, nil
}
