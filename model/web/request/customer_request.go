package request

type CustomerRegister struct {
	Email    string `json:"email" validate:"required,email,email_unique=customers" example:"dihanto@gmail.com"`
	Name     string `json:"name" validate:"required,min=5" example:"dihanto"`
	Password string `json:"password" validate:"required,min=8" example:"awesdkoire"`
}

type CustomerLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type CustomerUpdate struct {
	Name  string `json:"name" validate:"required,min=5"`
	Email string `json:"email" validate:"required,email"`
}

type CustomerDelete struct {
	Email string `json:"email" validate:"required,email"`
}
