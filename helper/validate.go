package helper

import (
	"context"

	"github.com/dihanto/go-toko/config"
	"github.com/dihanto/go-toko/exception"
	"github.com/go-playground/validator/v10"
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
		exception.ErrorHandler(nil, nil, err)
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		err = rows.Scan(&email)
		if err != nil {
			exception.ErrorHandler(nil, nil, err)
		}
		if value == email {
			return false
		}
	}

	return true

}
