// Package errs provides a structured approach to error handling within the application.
package errs

import (
	"fmt"
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
