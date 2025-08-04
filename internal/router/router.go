// Package router defines HTTP routing logic and middleware for the Projectify API.
package router

import (
	"net/http"

	customerhandlers "github.com/leonlonsdale/projectify/internal/router/customer"
	storage "github.com/leonlonsdale/projectify/internal/storage"
)

type Router struct {
	Customers *customerhandlers.CustomerHandler
}

func NewRouter(store *storage.Storage) *Router {
	return &Router{
		Customers: customerhandlers.NewCustomerHandler(store),
	}
}

func (r *Router) Mux() *http.ServeMux {
	m := http.NewServeMux()

	r.Customers.Register(m)

	return m
}
