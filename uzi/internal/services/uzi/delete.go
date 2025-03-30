package uzi

import (
	"context"

	"github.com/google/uuid"
)

func (s *service) DeleteUzi(ctx context.Context, id uuid.UUID) error {
	return s.dao.NewUziQuery(ctx).DeleteUzi(id)
}
