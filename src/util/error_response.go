package util

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
