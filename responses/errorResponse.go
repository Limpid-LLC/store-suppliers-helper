package responses

type ErrorResponse struct {
	Code    any    `json:"code"`
	Message string `json:"message"`
}
