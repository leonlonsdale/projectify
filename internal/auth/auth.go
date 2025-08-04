// Package auth exposes an Auth struct and utility functions for Argon2id
// password encryption and JWT methods.
package auth

type Auth struct{}

func NewAuth() *Auth {
	return &Auth{}
}
