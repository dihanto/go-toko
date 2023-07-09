package response

import (
	"time"

	"github.com/google/uuid"
)

type AddWallet struct {
	Id         int       `json:"id"`
	IdCustomer uuid.UUID `json:"id_customer"`
	Balance    int       `json:"balance"`
	CreatedAt  time.Time `json:"created_at"`
}

type GetWallet struct {
	Id         int       `json:"id"`
	IdCustomer uuid.UUID `json:"id_customer"`
	Balance    int       `json:"balance"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UpdateWallet struct {
	IdCustomer uuid.UUID `json:"id_customer"`
	Balance    int       `json:"balance"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
