package helper

type Login struct {
	EmailRequest string `json:"email_request" validate:"required"`
	EmailResult  string `json:"email_result" validate:"required,eqfield=EmailRequest"`
}
