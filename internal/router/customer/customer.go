// Package customerhandlers provides crud handler functions for customer actions.
package customerhandlers

import (
	"net/http"

	"github.com/leonlonsdale/projectify/internal/auth"
	"github.com/leonlonsdale/projectify/internal/storage"
	"github.com/leonlonsdale/projectify/pkg/utils/httputils"
)

type CustomerHandler struct {
	store *storage.Storage
	auth  *auth.Auth
}

func NewCustomerHandler(store *storage.Storage, auth *auth.Auth) *CustomerHandler {
	return &CustomerHandler{
		store: store,
		auth:  auth,
	}
}

func (h *CustomerHandler) Register(m *http.ServeMux) {
	m.Handle("POST /customer", httputils.Make(h.HandleCreateCustomer))
	m.Handle("GET /customer/{id}", httputils.Make(h.HandleGetCustomerByID))
	m.Handle("GET /customer", httputils.Make(h.HandleGetCustomerByEmail))
	m.Handle("PUT /customer/{id}", httputils.Make(h.HandleUpdateCustomer))
	m.Handle("DELETE /customer/{id}", httputils.Make(h.HandleDeleteCustomer))
}
