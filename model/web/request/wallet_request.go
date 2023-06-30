package request

import "github.com/google/uuid"

type AddWallet struct {
	IdCustomer uuid.UUID `json:"id_customer" validate:"required"`
	Balance    int       `json:"balance" validate:"required"`
}

type UpdateWallet struct {
	IdCustomer uuid.UUID `json:"id_customer" validate:"required"`
	Balance    int       `json:"balance" validate:"required"`
}
