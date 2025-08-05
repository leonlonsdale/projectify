package customerstore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/leonlonsdale/projectify/internal/errs"
	"github.com/leonlonsdale/projectify/pkg/utils/pgutils"
)

func (s *CustomerStorage) Delete(ctx context.Context, id uuid.UUID) error {

	query := `
		DELETE FROM customers
		WHERE id = $1;
	`

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		mappedErr := pgutils.MapPgError(err)
		if errors.Is(err, sql.ErrNoRows) {
			return errs.ErrUserNotFound
		}
		return mappedErr
	}

	return nil

}
