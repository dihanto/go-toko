package response

import "time"

type CustomerRegister struct {
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	RegisteredAt time.Time `json:"registered_at"`
}

type CustomerUpdate struct {
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	RegisteredAt time.Time `json:"registered_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
