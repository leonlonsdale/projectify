package errs

import (
	"errors"
	"fmt"
	"net/http"
)

type Kind string

const (
	KindBadRequest   Kind = "BadRequest"
	KindInternal     Kind = "InternalServerError"
	KindNotFound     Kind = "NotFound"
	KindUnauthorized Kind = "Unauthorized"
	KindForbidden    Kind = "Forbidden"
)

type Error struct {
	Err        error
	Kind       Kind
	StatusCode int
	Message    any
}

func (e *Error) Error() string {
	var msg string
	if m, ok := e.Message.(string); ok {
		msg = m
	} else {
		msg = fmt.Sprintf("%v", e.Message)
	}

	if e.Err != nil {
		return fmt.Sprintf("[%d] %s: %s (underlying error: %v)", e.StatusCode, e.Kind, msg, e.Err)
	}
	return fmt.Sprintf("[%d] %s: %s", e.StatusCode, e.Kind, msg)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func New(kind Kind, statusCode int, message string) error {
	return &Error{
		Err:        nil,
		Kind:       kind,
		StatusCode: statusCode,
		Message:    message,
	}
}

func Wrap(err error, kind Kind, statusCode int, message string) error {
	if err == nil {
		return nil
	}
	return &Error{
		Err:        err,
		Kind:       kind,
		StatusCode: statusCode,
		Message:    message,
	}
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if appErr, ok := err.(*Error); ok {
		return appErr.StatusCode
	}

	return http.StatusInternalServerError
}

func NewBadRequest(message string, err error) error {
	return Wrap(err, KindBadRequest, http.StatusBadRequest, message)
}

func NewInternalServerError(message string, err error) error {
	return Wrap(err, KindInternal, http.StatusInternalServerError, message)
}

func NewNotFound(message string, err error) error {
	return Wrap(err, KindNotFound, http.StatusNotFound, message)
}

func NewUnauthorized(message string, err error) error {
	return Wrap(err, KindUnauthorized, http.StatusUnauthorized, message)
}

func NewForbidden(message string, err error) error {
	return Wrap(err, KindForbidden, http.StatusForbidden, message)
}

func NewValidationErrors(validationErrs map[string]string) error {
	if len(validationErrs) == 0 {
		return nil
	}
	return &Error{
		Err:        nil,
		Kind:       KindBadRequest,
		StatusCode: http.StatusBadRequest,
		Message:    validationErrs,
	}
}

type Response struct {
	Status int      `json:"status"`
	Error  *Details `json:"error,omitempty"`
}

type Details struct {
	Kind    Kind `json:"kind"`
	Message any  `json:"message"`
}

func ToJSON(err error) Response {
	var appErr *Error
	if !errors.As(err, &appErr) {
		return Response{
			Status: http.StatusInternalServerError,
			Error: &Details{
				Kind: KindInternal,
				Message: map[string]string{
					"error": "An unexpected internal error occurred",
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
