package response

type ErrorResponse struct {
	Code    int `json:"code"`
	Message any `json:"message"`
}
