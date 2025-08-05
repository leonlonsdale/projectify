package errs

import "errors"

var (
	ErrInvalidCredentials        = errors.New("username or password incorrect")
	ErrEmailAlreadyExists        = errors.New("email already exists")
	ErrUserNotFound              = errors.New("user not found")
	ErrorJWTInvalidSigningMethod = errors.New("invalid signing method")
	ErrorJWTInvalidToken         = errors.New("invalid jwt token")
	ErrorJWTInvalidClaimsType    = errors.New("invalid claims type")
	ErrorJWTTokenHasExpired      = errors.New("jwt token has expired")
	ErrorNoBearerToken           = errors.New("no bearer token in headers")
)
