package response

import (
	"time"

	"github.com/google/uuid"
)

type AddProduct struct {
	Id        int       `json:"id"`
	IdSeller  uuid.UUID `json:"id_seller"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

type GetProduct struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type FindById struct {
	Id        int       `json:"id"`
	IdSeller  uuid.UUID `json:"id_seller"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateProduct struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FindByName struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}
