package entity

import "github.com/google/uuid"

type Order struct {
	Id         int
	IdProduct  int
	IdCustomer uuid.UUID
	Quantity   int
}
