package node_segment

import (
	"context"

	"github.com/google/uuid"

	domain "gateway/internal/domain/uzi"
)

func (s *service) GetNodeWithSegmentsByImageID(ctx context.Context, imageID uuid.UUID) ([]domain.Node, []domain.Segment, error) {
	nodes, segments, err := s.adapters.Uzi.GetNodesWithSegmentsByImageId(ctx, imageID)
	if err != nil {
		return nil, nil, err
	}
	return nodes, segments, nil
}
