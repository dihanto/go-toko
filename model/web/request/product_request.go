package request

import "github.com/google/uuid"

type AddProduct struct {
	IdSeller uuid.UUID `json:"id_seller"`
	Name     string    `json:"name"`
	Price    int       `json:"price"`
	Quantity int       `json:"quantity"`
}

type UpdateProduct struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}
