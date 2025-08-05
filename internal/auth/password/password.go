// Package password provides methods for hashing and validating passwords
package password

import (
	"github.com/alexedwards/argon2id"
	"github.com/leonlonsdale/projectify/internal/errs"
)

var params = &argon2id.Params{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

type A2idpassword struct{}

func NewA2idpassword() *A2idpassword {
	return &A2idpassword{}
}

func (p *A2idpassword) Hash(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, params)

	if err != nil {
		return "", err
	}

	return hash, nil
}

func (p *A2idpassword) Validate(password, hash string) error {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return err
	}

	if !match {
		return errs.ErrInvalidCredentials
	}

	return nil
}
