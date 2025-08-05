package customerhandlers

import (
	"errors"
	"net/http"

	"github.com/leonlonsdale/projectify/internal/errs"
	"github.com/leonlonsdale/projectify/internal/models"
	"github.com/leonlonsdale/projectify/pkg/utils/jsonutils"
)

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
