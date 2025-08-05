// Package customerstore exposes database manipulation methods for customers.
package customerstore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/leonlonsdale/projectify/internal/errs"
	"github.com/leonlonsdale/projectify/internal/models"
	"github.com/leonlonsdale/projectify/pkg/utils/pgutils"
)

type CustomerStorage struct {
	db *sql.DB
}

func NewCustomerStorage(db *sql.DB) *CustomerStorage {
	return &CustomerStorage{
		db: db,
	}
}

func (s *CustomerStorage) Create(ctx context.Context, data models.CustomerRegistration) (*models.CustomerSafe, error) {
	query := `
		INSERT INTO customers (id, created_at, updated_at, name, email, password)
		VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3)
		RETURNING id, created_at, updated_at, name, email;
	`

	row := s.db.QueryRowContext(ctx, query, data.Name, data.Email, data.Password)

	var u models.CustomerSafe
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

func (s *CustomerStorage) GetByID(ctx context.Context, id uuid.UUID) (models.CustomerSafe, error) {

	return models.CustomerSafe{}, nil
}

func (s *CustomerStorage) GetByEmail(ctx context.Context, email string) (models.CustomerSafe, error) {
	return models.CustomerSafe{}, nil
}

func (s *CustomerStorage) Update(ctx context.Context, data models.CustomerUpdate, id uuid.UUID) (*models.CustomerSafe, error) {
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

	var u models.CustomerSafe
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

func (s *CustomerStorage) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
