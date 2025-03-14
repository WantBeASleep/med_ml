package node

import (
	"context"

	"uzi/internal/domain"
	nodeEntity "uzi/internal/repository/node/entity"

	"github.com/google/uuid"
)

func (s *service) GetNodesByUziID(ctx context.Context, id uuid.UUID) ([]domain.Node, error) {
	nodes, err := s.dao.NewNodeQuery(ctx).GetNodesByUziID(id)
	if err != nil {
		return nil, err
	}

	return nodeEntity.Node{}.SliceToDomain(nodes), nil
}
