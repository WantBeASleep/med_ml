package node

import (
	"context"

	"github.com/google/uuid"

	domain "composition-api/internal/domain/uzi"
)

func (s *service) GetNodesByUziID(ctx context.Context, uziID uuid.UUID) ([]domain.Node, error) {
	nodes, err := s.adapters.Uzi.GetNodesByUziId(ctx, uziID)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}
