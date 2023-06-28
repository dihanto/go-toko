package response

import "github.com/google/uuid"

type AddWallet struct {
	Id         int       `json:"id"`
	IdCustomer uuid.UUID `json:"id_customer"`
	Balance    int       `json:"balance"`
}

type GetWallet struct {
	Id      int `json:"id"`
	Balance int `json:"balance"`
}

type UpdateWallet struct {
	Id      int `json:"id"`
	Balance int `json:"balance"`
}
