package errs

import "net/http"

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
