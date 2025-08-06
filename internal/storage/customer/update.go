package customerstore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/leonlonsdale/projectify/internal/api"
	"github.com/leonlonsdale/projectify/internal/errs"
	"github.com/leonlonsdale/projectify/pkg/utils/pgutils"
)

func (s *CustomerStorage) Update(ctx context.Context, data api.CustomerUpdate, id uuid.UUID) (*api.CustomerSafe, error) {
	query := `
		UPDATE customers
		SET
			name = COALESCE($1, name),
			email = COALESCE($2, name),
			updated_at = NOW()
		WHERE id = $3
		RETURNING id, created_at, updated_at, name, email;
	`

	row := s.db.QueryRowContext(ctx, query, data.Name, data.Email, id)

	var u api.CustomerSafe
	err := row.Scan(
		&u.ID,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Name,
		&u.Email,
	)
	if err != nil {
		mappedErr := pgutils.MapPgError(err)
		if errors.Is(mappedErr, sql.ErrNoRows) {
			return nil, errs.ErrUserNotFound
		}

		return nil, err
	}

	return &u, nil
}
