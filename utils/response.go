package utils

type Response[T any] struct {
	Data         *T      `json:"data"`
	Error        bool    `json:"error"`
	ErrorCode    *int    `json:"error_code"`
	ErrorMessage *string `json:"error_message"`
}
