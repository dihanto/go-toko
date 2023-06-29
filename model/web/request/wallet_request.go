package request

import "github.com/google/uuid"

type AddWallet struct {
	IdCustomer uuid.UUID `json:"id_customer"`
	Balance    int       `json:"balance"`
}

type UpdateWallet struct {
	IdCustomer uuid.UUID `json:"id_customer"`
	Balance    int       `json:"balance"`
}
