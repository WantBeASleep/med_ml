package node

import (
	"context"

	"github.com/google/uuid"
)

func (s *service) DeleteNode(ctx context.Context, id uuid.UUID) error {
	return s.adapters.Uzi.DeleteNode(ctx, id)
}
