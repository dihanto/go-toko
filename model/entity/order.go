package entity

import "github.com/google/uuid"

type Order struct {
	Id         int
	IdProduct  int
	IdCustomer uuid.UUID
	Quantity   int
	OrderedAt  int32
	UpdatedAt  int32
	DeletedAt  int32
}
