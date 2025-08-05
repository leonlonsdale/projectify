// Package httputils exposes helper utility functions for http.
package httputils

import (
	"log/slog"
	"net/http"

	"github.com/leonlonsdale/projectify/internal/errs"
	"github.com/leonlonsdale/projectify/pkg/utils/jsonutils"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

func Make(f HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			slog.Error("API Error", "err", err.Error(), "path", r.URL.Path)
			jsonResponse := errs.ErrToJSON(err)
			_ = jsonutils.WriteJSON(w, jsonResponse.Status, jsonResponse)
		}
	}
}
