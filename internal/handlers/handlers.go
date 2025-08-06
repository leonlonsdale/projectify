// Package handlers defines HTTP routing logic and middleware for the Projectify API.
package handlers

import (
	"net/http"

	"github.com/leonlonsdale/projectify/internal/auth"
	"github.com/leonlonsdale/projectify/internal/config"
	customerhandlers "github.com/leonlonsdale/projectify/internal/handlers/customer"
	"github.com/leonlonsdale/projectify/internal/storage"
)

type Handlers struct {
	Customers *customerhandlers.CustomerHandler
}

func NewRouter(store *storage.Storage, cfg *config.Config, auth *auth.Auth) *Handlers {
	return &Handlers{
		Customers: customerhandlers.NewCustomerHandler(store, auth),
	}
}

func (r *Handlers) Mux() *http.ServeMux {
	m := http.NewServeMux()

	r.Customers.Register(m)

	fs := http.FileServer(http.Dir("./static"))
	m.Handle("/static/", http.StripPrefix("/static/", fs))

	return m
}
