package entity

import "github.com/google/uuid"

type Product struct {
	Id        int
	IdSeller  uuid.UUID
	Name      string
	Price     int
	Quantity  int
	CreatedAt int32
	UpdatedAt int32
	DeletedAt int32
}
