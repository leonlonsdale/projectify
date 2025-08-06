package customerstore

import (
	"context"

	"github.com/leonlonsdale/projectify/internal/api"
	"github.com/leonlonsdale/projectify/internal/errs"
	"github.com/leonlonsdale/projectify/pkg/utils/pgutils"
)

func (s *CustomerStorage) Create(ctx context.Context, data api.CustomerRegistration) (*api.CustomerSafe, error) {
	query := `
		INSERT INTO customers (id, created_at, updated_at, name, email, password)
		VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3)
		RETURNING id, created_at, updated_at, name, email;
	`

	row := s.db.QueryRowContext(ctx, query, data.Name, data.Email, data.Password)

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

		if mappedErr.Error() == "unique_constraint_violation" {
			return nil, errs.ErrEmailAlreadyExists
		}

		return nil, mappedErr
	}

	return &u, nil
}
