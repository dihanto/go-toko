package request

import "github.com/google/uuid"

type AddOrder struct {
	IdProduct  int       `json:"id_product" validate:"required"`
	IdCustomer uuid.UUID `json:"id_customer" validate:"required"`
	Quantity   int       `json:"quantity" validate:"required"`
}
