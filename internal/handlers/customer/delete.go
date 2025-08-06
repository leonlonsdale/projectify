package customerhandlers

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/leonlonsdale/projectify/internal/errs"
	"github.com/leonlonsdale/projectify/pkg/utils/httputils"
)

func (h *CustomerHandler) HandleDeleteCustomer(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	userIDFromJWT, ok := h.auth.UserIDFromContext(ctx)
	if !ok {
		return errs.NewUnauthorized("user not logged in", nil)
	}

	userIDFromPath := r.PathValue("id")
	parsedUserIDFromPath, err := uuid.Parse(userIDFromPath)
	if err != nil {
		return errs.NewBadRequest("invalid user id", nil)
	}

	if userIDFromJWT != parsedUserIDFromPath {
		return errs.NewForbidden("a user may only delete their own profile", nil)
	}

	if err := h.store.Customers.Delete(ctx, parsedUserIDFromPath); err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			return errs.NewNotFound("user not found for this id", err)
		}
		return errs.NewInternalServerError("an unexpected internal server error occurred", err)
	}

	if err := httputils.WriteSuccessJSON(w, http.StatusNoContent, "", nil); err != nil {
		return errs.NewInternalServerError("failed to write response to json", err)
	}

	return nil
}
