// Package handlers...

package handlers

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/leonlonsdale/projectify/internal/errs"
	utils "github.com/leonlonsdale/projectify/pkg/util"
)

type Handlers struct {
	Customers *CustomerHandler
}

func NewHandlers(db *sql.DB) *Handlers {
	return &Handlers{
		Customers: NewCustomerHandler(db),
	}
}

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

func Make(f HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			slog.Error("HTTP API Error", "err", err.Error(), "path", r.URL.Path)
			jsonResponse := errs.ToJSON(err)
			_ = utils.WriteJSON(w, jsonResponse.Status, jsonResponse)
		}
	}
}
