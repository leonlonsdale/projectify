package httputils

import (
	"log/slog"
	"net/http"

	"github.com/leonlonsdale/projectify/internal/api"
	"github.com/leonlonsdale/projectify/internal/errs"
	"github.com/leonlonsdale/projectify/pkg/utils/jsonutils"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

func Make(f HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); r != nil {
				slog.Error("API Panic", "panic", p, "path", r.URL.Path)

				err := errs.NewInternalServerError("an unexpected internal server error occurred", nil)
				_ = WriteErrorJSON(w, err)
			}
		}()
		if err := f(w, r); err != nil {
			slog.Error("API Error", "err", err.Error(), "path", r.URL.Path)
			_ = WriteErrorJSON(w, err)
		}
	}
}

func WriteResponseJSON(w http.ResponseWriter, code int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	jsonBytes, err := jsonutils.EncodeJSON(data)
	if err != nil {
		slog.Error("failed to marshal JSON response", "error", err)
		return err
	}

	_, err = w.Write(jsonBytes)
	return err
}

func WriteSuccessJSON(w http.ResponseWriter, code int, key string, data any) error {
	payloadData := map[string]any{
		key: data,
	}

	payload := api.Payload[any]{
		Data: payloadData,
	}

	successResp := api.NewSuccessResponse(code, payload)

	return WriteResponseJSON(w, code, successResp)
}

func WriteErrorJSON(w http.ResponseWriter, err error) error {
	errResp := errs.ErrToJSON(err)
	return WriteResponseJSON(w, errResp.Status, errResp)
}
