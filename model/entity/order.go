package entity

import "github.com/google/uuid"

type Order struct {
	Id         int
	IdCustomer uuid.UUID
	OrderedAt  int32
	UpdatedAt  int32
	DeletedAt  int32
}
