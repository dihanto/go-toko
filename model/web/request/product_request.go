package request

import "github.com/google/uuid"

type AddProduct struct {
	IdSeller uuid.UUID `json:"id_seller" validate:"required"`
	Name     string    `json:"name" validate:"required"`
	Price    int       `json:"price" validate:"required"`
	Quantity int       `json:"quantity" validate:"required"`
}

type UpdateProduct struct {
	Id       int    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Price    int    `json:"price" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}

type AddProductToWishlist struct {
	ProductId  int       `json:"product_id" validate:"required"`
	CustomerId uuid.UUID `json:"user_id" validate:"required"`
}

type DeleteProductFromWishlist struct {
	ProductId  int       `json:"product_id" validate:"required"`
	CustomerId uuid.UUID `json:"user_id" validate:"required"`
}
