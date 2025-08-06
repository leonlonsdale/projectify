package customerstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/leonlonsdale/projectify/internal/api"
	"github.com/leonlonsdale/projectify/internal/errs"
	"github.com/leonlonsdale/projectify/pkg/utils/pgutils"
)

func (s *CustomerStorage) get(ctx context.Context, column string, value any) (*api.CustomerSafe, error) {

	validColumns := map[string]struct{}{
		"id":    {},
		"email": {},
	}

	if _, ok := validColumns[column]; !ok {
		return nil, fmt.Errorf("invalid column provided: %s", column)
	}

	query := `
		SELECT id, created_at, updated_at, name, email
		FROM users
		WHERE ` + column + ` = $1;
	`

	row := s.db.QueryRowContext(ctx, query, value)

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

		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrUserNotFound
		}

		return nil, mappedErr
	}

	return &u, nil
}

func (s *CustomerStorage) GetByID(ctx context.Context, id uuid.UUID) (*api.CustomerSafe, error) {
	return s.get(ctx, "id", id)
}

func (s *CustomerStorage) GetByEmail(ctx context.Context, email string) (*api.CustomerSafe, error) {
	return s.get(ctx, "email", email)
}
