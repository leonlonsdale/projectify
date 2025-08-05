// Package customerhandlers provides crud handler functions for customer actions.
package customerhandlers

import (
	"errors"
	"net/http"

	"github.com/leonlonsdale/projectify/internal/auth"
	"github.com/leonlonsdale/projectify/internal/errs"
	"github.com/leonlonsdale/projectify/internal/models"
	"github.com/leonlonsdale/projectify/internal/storage"
	"github.com/leonlonsdale/projectify/pkg/utils/httputils"
	"github.com/leonlonsdale/projectify/pkg/utils/jsonutils"
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

func (h *CustomerHandler) HandleCreateCustomer(w http.ResponseWriter, r *http.Request) error {

	ctx := r.Context()
	var params models.CustomerRegistration
	if err := jsonutils.DecodeJSON(r.Body, &params); err != nil {
		return errs.NewBadRequest("malformed request body", err)
	}

	if validationErrs := params.Validate(); validationErrs != nil {
		return errs.NewValidationErrors(validationErrs)
	}

	hashedPass, err := h.auth.Password.Hash(params.Password)
	if err != nil {
		return errs.NewInternalServerError("error hashing password", err)
	}

	params.Password = hashedPass

	user, err := h.store.Customers.Create(ctx, params)
	if err != nil {
		if errors.Is(err, errs.ErrEmailAlreadyExists) {
			return errs.NewBadRequest("Could not create customer", err)
		}
		return errs.NewInternalServerError("failed to create customer", err)
	}

	safeUser := models.CustomerSafe{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     user.Email,
	}

	if err := jsonutils.WriteJSON(w, http.StatusCreated, safeUser); err != nil {
		return errs.NewInternalServerError("error writing json", err)
	}

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
