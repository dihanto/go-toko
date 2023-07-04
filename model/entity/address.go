package entity

import "github.com/google/uuid"

type Address struct {
	Code       string
	IdCustomer uuid.UUID
	Street     string
	City       string
	Province   string
	CreatedAt  int32
	UpdatedAt  int32
	DeletedAt  int32
}
