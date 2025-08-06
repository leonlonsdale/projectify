package api

type ErrorResponse struct {
	Status int           `json:"status"`
	Error  *ErrorDetails `json:"error,omitempty"`
}

type ErrorDetails struct {
	Kind    string `json:"kind"`
	Message any    `json:"message"`
}

type SuccessResponse[T any] struct {
	Status int `json:"status"`
	Data   T   `json:"data,omitempty"`
}

type Payload[T any] struct {
	Data T `json:"data"`
}

func NewErrorResponse(status int, kind string, message any) *ErrorResponse {
	return &ErrorResponse{
		Status: status,
		Error: &ErrorDetails{
			Kind:    kind,
			Message: message,
		},
	}
}

func NewSuccessResponse[T any](status int, data T) *SuccessResponse[T] {
	return &SuccessResponse[T]{
		Status: status,
		Data:   data,
	}
}
