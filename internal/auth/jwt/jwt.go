// Package jwt is part of the auth layer and provdes JWT functionality.
package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/leonlonsdale/projectify/internal/config"
	"github.com/leonlonsdale/projectify/internal/errs"
)

type TokenType string

const (
	TokenTypeAccess TokenType = "projectify"
)

type JWT struct {
	CFG *config.Config
}

func NewJWT(cfg *config.Config) *JWT {
	return &JWT{
		CFG: cfg,
	}
}

func (j *JWT) Make(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {

	now := time.Now().UTC()

	claims := jwt.RegisteredClaims{
		Issuer:    string(TokenTypeAccess),
		Subject:   userID.String(),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(expiresIn)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", errs.NewInternalServerError("error signing jwt token", err)
	}

	return tokenString, nil
}

func (j *JWT) Validate(tokenString, tokenSecret string) (uuid.UUID, error) {
	claims := &jwt.RegisteredClaims{}
	keyfunc := func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return uuid.Nil, errs.ErrorJWTInvalidSigningMethod
		}
		return []byte(tokenSecret), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, keyfunc)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return uuid.Nil, errs.ErrorJWTInvalidToken
		}
		return uuid.Nil, errs.ErrorJWTInvalidToken
	}

	if !token.Valid {
		return uuid.Nil, errs.ErrorJWTInvalidToken
	}

	subject, err := claims.GetSubject()
	if err != nil {
		return uuid.Nil, errs.ErrorJWTInvalidClaimsType
	}

	issuer, err := claims.GetIssuer()
	if err != nil || issuer != string(TokenTypeAccess) {
		return uuid.Nil, errs.ErrorJWTInvalidClaimsType
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now().UTC()) {
		return uuid.Nil, errs.ErrorJWTTokenHasExpired
	}

	userID, err := uuid.Parse(subject)
	if err != nil {
		return uuid.Nil, errs.NewInternalServerError("malformed user id in token claims", err)
	}

	return userID, nil
}
