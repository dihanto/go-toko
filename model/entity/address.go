package entity

import "github.com/google/uuid"

type Address struct {
	Code       string
	IdCustomer uuid.UUID
	Street     string
	City       string
	Province   string
}
