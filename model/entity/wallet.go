package entity

import "github.com/google/uuid"

type Wallet struct {
	Id         int
	IdCustomer uuid.UUID
	Balance    int
	CreatedAt  int32
	UpdatedAt  int32
	DeletedAt  int32
}
