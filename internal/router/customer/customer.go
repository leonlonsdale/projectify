// Package customerhandlers provides crud handler functions for customer actions.
package customerhandlers

import (
	"net/http"

	"github.com/leonlonsdale/projectify/internal/storage"
	utils "github.com/leonlonsdale/projectify/pkg/util"
)

type CustomerHandler struct {
	store *storage.Storage
}

func NewCustomerHandler(store *storage.Storage) *CustomerHandler {
	return &CustomerHandler{store}
}

func (h *CustomerHandler) Register(m *http.ServeMux) {
	m.Handle("POST /customer", utils.Make(h.HandleCreateCustomer))
	m.Handle("GET /customer/{id}", utils.Make(h.HandleGetCustomerByID))
	m.Handle("GET /customer", utils.Make(h.HandleGetCustomerByEmail))
	m.Handle("PUT /customer/{id}", utils.Make(h.HandleUpdateCustomer))
	m.Handle("DELETE /customer/{id}", utils.Make(h.HandleDeleteCustomer))
}

func (h *CustomerHandler) HandleCreateCustomer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *CustomerHandler) HandleGetCustomerByID(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *CustomerHandler) HandleGetCustomerByEmail(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *CustomerHandler) HandleUpdateCustomer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *CustomerHandler) HandleDeleteCustomer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
