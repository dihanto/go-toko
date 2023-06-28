package entity

import "github.com/google/uuid"

type Wallet struct {
	Id         int
	IdCustomer uuid.UUID
	Balance    int
}
