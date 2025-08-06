package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/leonlonsdale/projectify/internal/errs"
	"github.com/leonlonsdale/projectify/pkg/utils/httputils"
)

type contextKey string

const UserIDKey contextKey = "userID"

func (a *Auth) Protect(next httputils.HTTPHandler) httputils.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		tokenString, err := a.JWT.GetBearerToken(r.Header)
		if err != nil {
			return errs.NewUnauthorized("missing or malformed authorization header", err)
		}

		userID, err := a.JWT.Validate(tokenString, a.JWT.CFG.JWTSecret)
		if err != nil {
			if errors.Is(err, errs.ErrorJWTInvalidToken) || errors.Is(err, errs.ErrorJWTInvalidClaimsType) || errors.Is(err, errs.ErrorJWTInvalidSigningMethod) {

				return errs.NewUnauthorized("invalid or expired token", err)
			}
			return errs.NewInternalServerError("failed to validatae token", err)
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		return next(w, r.WithContext(ctx))
	}
}

func (a *Auth) UserIDFromContext(ctx context.Context) (uuid.UUID, bool) {
	userID, ok := ctx.Value(UserIDKey).(uuid.UUID)
	if !ok {
		return uuid.Nil, false
	}

	return userID, true
}
