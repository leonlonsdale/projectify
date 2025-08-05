package customerstore

import (
	"context"

	"github.com/google/uuid"
)

func (s *CustomerStorage) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
