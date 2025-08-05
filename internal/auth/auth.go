// Package auth exposes an Auth layer which exposes password and jwt functionality
package auth

import (
	"github.com/leonlonsdale/projectify/internal/auth/jwt"
	"github.com/leonlonsdale/projectify/internal/auth/password"
	"github.com/leonlonsdale/projectify/internal/config"
)

type Auth struct {
	Password *password.A2idpassword
	JWT      *jwt.JWT
}

func NewAuth(cfg *config.Config) *Auth {
	return &Auth{
		Password: password.NewA2idpassword(),
		JWT:      jwt.NewJWT(cfg),
	}
}
