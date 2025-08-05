package customerstore

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonlonsdale/projectify/internal/models"
)

func (s *CustomerStorage) GetByID(ctx context.Context, id uuid.UUID) (models.CustomerSafe, error) {

	return models.CustomerSafe{}, nil
}

func (s *CustomerStorage) GetByEmail(ctx context.Context, email string) (models.CustomerSafe, error) {
	return models.CustomerSafe{}, nil
}
