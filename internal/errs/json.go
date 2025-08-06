package errs

import (
	"errors"
	"net/http"

	"github.com/leonlonsdale/projectify/internal/api"
)

const InternalErrorMsg = "an unexpected internal server error occurred"

func ErrToJSON(err error) *api.ErrorResponse {
	var appErr *Error
	if !errors.As(err, &appErr) {
		return api.NewErrorResponse(
			http.StatusInternalServerError,
			string(KindInternal),
			InternalErrorMsg,
		)
	}

	if appErr.Kind == KindInternal {
		return api.NewErrorResponse(
			http.StatusInternalServerError,
			string(KindInternal),
			InternalErrorMsg,
		)
	}

	return api.NewErrorResponse(
		appErr.StatusCode,
		string(appErr.Kind),
		appErr.Message,
	)
}
