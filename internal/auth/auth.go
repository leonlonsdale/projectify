// Package auth exposes an Auth layer which exposes password and jwt functionality
package auth

import "github.com/leonlonsdale/projectify/internal/auth/password"

type Auth struct {
	Password *password.A2idpassword
}

func NewAuth() *Auth {
	return &Auth{
		Password: password.NewA2idpassword(),
	}
}
