package request

type CustomerRegister struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CustomerLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomerUpdate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CustomerDelete struct {
	Email string `json:"email"`
}
