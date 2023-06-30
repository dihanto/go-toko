package request

type SellerRegister struct {
	Email    string `json:"email" validate:"required,email,email_unique=sellers"`
	Name     string `json:"name" validate:"required,min=5"`
	Password string `json:"password" validate:"required,min=8"`
}

type SellerLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type SellerUpdate struct {
	Name  string `json:"name" validate:"required,min=5"`
	Email string `json:"email" validate:"required,email"`
}

type SellerDelete struct {
	Email string `json:"email" validate:"required,email"`
}
