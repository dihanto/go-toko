package request

type SellerRegister struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SellerLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SellerUpdate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SellerDelete struct {
	Email string `json:"email"`
}
