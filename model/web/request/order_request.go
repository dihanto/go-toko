package request

import "github.com/google/uuid"

type AddOrder struct {
	IdProduct  int       `json:"id_product"`
	IdCustomer uuid.UUID `json:"id_customer"`
	Quantity   int       `json:"quantity"`
}
