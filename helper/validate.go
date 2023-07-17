package helper

import (
	"context"
	"log"

	"github.com/dihanto/go-toko/config"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func ValdateEmailUnique(field validator.FieldLevel) bool {
	value := field.Field().Interface().(string)
	user := field.Param()

	conn := config.InitDatabaseConnection()
	defer conn.Close()

	ctx := context.Background()

	query := "SELECT email FROM " + user
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		err = rows.Scan(&email)
		if err != nil {
			log.Println(err)
		}
		if value == email {
			return false
		}
	}

	return true

}

func ValidateUserOnlyHaveOneWallet(field validator.FieldLevel) bool {
	value := field.Field().Interface().(uuid.UUID)

	conn := config.InitDatabaseConnection()
	defer conn.Close()

	ctx := context.Background()

	query := "SELECT id_customer FROM wallets"
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id uuid.UUID
		err = rows.Scan(&id)
		if err != nil {
			log.Println(err)
		}
		if id == value {
			return false
		}
	}
	return true
}

func ValidateUserHaveOneWishlistInOneProduct(field validator.FieldLevel) bool {
	value := field.Field().Interface().(uuid.UUID)
	productId := field.Param()

	conn := config.InitDatabaseConnection()
	defer conn.Close()

	ctx := context.Background()

	query := "SELECT customer_id FROM wishlist_details WHERE product_id=$1"
	rows, err := conn.QueryContext(ctx, query, productId)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id uuid.UUID
		err = rows.Scan(&id)
		if err != nil {
			log.Println(err)
		}

		if value == id {
			return false
		}
	}

	return true
}
