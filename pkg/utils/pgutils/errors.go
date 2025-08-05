// Package pgutils provides func to map over postgres errors
package pgutils

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// MapPgError takes a raw database error and maps it to a canonical,
// domain-agnostic error.
func MapPgError(err error) error {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		case "23505":
			return errors.New("unique_constraint_violation")
		case "23503":
			return errors.New("foreign_key_violation")
		case "02000":
			return sql.ErrNoRows
		}
	}

	if errors.Is(err, sql.ErrNoRows) {
		return sql.ErrNoRows
	}

	return err
}
