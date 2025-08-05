package customerhandlers

import (
	"net/http"

	"github.com/leonlonsdale/projectify/internal/errs"
	"github.com/leonlonsdale/projectify/internal/models"
	"github.com/leonlonsdale/projectify/pkg/utils/jsonutils"
)

func (h *CustomerHandler) HandleUpdateCustomer(w http.ResponseWriter, r *http.Request) error {
	// ctx := r.Context()
	var params models.CustomerUpdate

	// TODO: Get userID from JWT
	/*
		userID, ok := h.auth.UserIDFromContext(r.Context())
		if !ok || userID == uuid.Nil {
			return errs.NewUnauthorized("user not logged in", nil)
		}
	*/
	if err := jsonutils.DecodeJSON(r.Body, &params); err != nil {
		return errs.NewBadRequest("malformed request body", err)
	}

	// TODO: Store user changes
	/*
		user, err := h.store.Customers.Update(ctx, params, userID)
		if err != nil {
			if errors.Is(err, errs.ErrUserNotFound) {
				return errs.NewNotFound("user not found", err)
			}

			return errs.NewInternalServerError("error updating user", err)
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
	*/

	return nil
}
