// Package router defines HTTP routing logic and middleware for the Projectify API.
package router

import (
	"net/http"

	"github.com/leonlonsdale/projectify/internal/auth"
	"github.com/leonlonsdale/projectify/internal/config"
	customerhandlers "github.com/leonlonsdale/projectify/internal/router/customer"
	storage "github.com/leonlonsdale/projectify/internal/storage"
)

type Router struct {
	Customers *customerhandlers.CustomerHandler
}

func NewRouter(store *storage.Storage, cfg *config.Config, auth *auth.Auth) *Router {
	return &Router{
		Customers: customerhandlers.NewCustomerHandler(store, auth),
	}
}

func (r *Router) Mux() *http.ServeMux {
	m := http.NewServeMux()

	r.Customers.Register(m)

	fs := http.FileServer(http.Dir("./static"))
	m.Handle("/static/", http.StripPrefix("/static/", fs))

	return m
}
