package errs

import "errors"

var (
	ErrInvalidCredentials = errors.New("username or password incorrect")
	ErrEmailAlreadyExists = errors.New("email already exists")
)
