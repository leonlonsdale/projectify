// Package customerstore exposes database manipulation methods for customers.
package customerstore

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/leonlonsdale/projectify/pkg/pgutils"
)

type CustomerRecord struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

type CustomerRegistration struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomerUpdate struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type CustomerSafe struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
}

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
)

type CustomerStorage struct {
	db *sql.DB
}

func NewCustomerStorage(db *sql.DB) *CustomerStorage {
	return &CustomerStorage{
		db: db,
	}
}

func (s *CustomerStorage) Create(ctx context.Context, data CustomerRegistration) (*CustomerSafe, error) {
	query := `
		INSERT INTO customers (id, created_at, updated_at, name, email, password)
		VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3)
		RETURNING id, created_at, updated_at, name, email;
	`

	row := s.db.QueryRowContext(ctx, query, data.Name, data.Email, data.Password)

	var u CustomerSafe
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
			return nil, ErrEmailAlreadyExists
		}

		return nil, mappedErr
	}

	return &u, nil
}

func (s *CustomerStorage) GetByID(ctx context.Context, id uuid.UUID) (CustomerSafe, error) {
	return CustomerSafe{}, nil
}

func (s *CustomerStorage) GetByEmail(ctx context.Context, email string) (CustomerSafe, error) {
	return CustomerSafe{}, nil
}

func (s *CustomerStorage) Update(ctx context.Context, data CustomerUpdate, id uuid.UUID) (CustomerSafe, error) {
	return CustomerSafe{}, nil
}

func (s *CustomerStorage) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
