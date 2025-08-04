// Package utils exposes helper utility functions.
package utils

import (
	"log/slog"
	"net/http"

	"github.com/leonlonsdale/projectify/internal/errs"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

func Make(f HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			slog.Error("API Error", "err", err.Error(), "path", r.URL.Path)
			jsonResponse := errs.ToJSON(err)
			_ = WriteJSON(w, jsonResponse.Status, jsonResponse)
		}
	}
}
