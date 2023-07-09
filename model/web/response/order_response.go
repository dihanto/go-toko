package response

import (
	"time"

	"github.com/google/uuid"
)

type AddOrder struct {
	Id         int       `json:"id"`
	IdProduct  int       `json:"id_product"`
	IdCustomer uuid.UUID `json:"id_customer"`
	Quantity   int       `json:"quantity"`
	OrderedAt  time.Time `json:"ordered_at"`
}

type FindOrder struct {
	Id         int           `json:"id"`
	IdProduct  int           `json:"id_product"`
	IdCustomer uuid.UUID     `json:"id_customer"`
	Quantity   int           `json:"quantity"`
	OrderedAt  time.Time     `json:"ordered_at"`
	TotalPrice int           `json:"total_price"`
	Product    ProductOrder  `json:"product"`
	Customer   CustomerOrder `json:"customer"`
}

type ProductOrder struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CustomerOrder struct {
	Name string `json:"name"`
}
