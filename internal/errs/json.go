package errs

import (
	"errors"
	"net/http"
)

type Response struct {
	Status int      `json:"status"`
	Error  *Details `json:"error,omitempty"`
}

type Details struct {
	Kind    Kind `json:"kind"`
	Message any  `json:"message"`
}

func ErrToJSON(err error) Response {
	var appErr *Error
	if !errors.As(err, &appErr) {
		return Response{
			Status: http.StatusInternalServerError,
			Error: &Details{
				Kind: KindInternal,
				Message: map[string]string{
					"error": "an unexpected internal error occurred",
				},
			},
		}
	}

	if appErr.Kind == KindInternal {
		return Response{
			Status: http.StatusInternalServerError,
			Error: &Details{
				Kind: KindInternal,
				Message: map[string]string{
					"error": "an unexpected internal error occurred",
				},
			},
		}
	}

	return Response{
		Status: appErr.StatusCode,
		Error: &Details{
			Kind:    appErr.Kind,
			Message: appErr.Message,
		},
	}
}
