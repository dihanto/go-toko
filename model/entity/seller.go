package entity

import "github.com/google/uuid"

type Seller struct {
	Id           uuid.UUID
	Email        string
	Name         string
	Password     string
	RegisteredAt int32
	UpdatedAt    int32
	DeletedAt    int32
}
