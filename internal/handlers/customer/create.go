package customerhandlers

import (
	"errors"
	"net/http"

	"github.com/leonlonsdale/projectify/internal/api"
	"github.com/leonlonsdale/projectify/internal/errs"
	"github.com/leonlonsdale/projectify/pkg/utils/httputils"
	"github.com/leonlonsdale/projectify/pkg/utils/jsonutils"
)

func (h *CustomerHandler) HandleCreateCustomer(w http.ResponseWriter, r *http.Request) error {

	ctx := r.Context()
	var params api.CustomerRegistration
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

	if err := httputils.WriteSuccessJSON(w, http.StatusCreated, "user", user); err != nil {
		return errs.NewInternalServerError("error writing json", err)
	}

	return nil
}
