package image_segment_node

import (
	"context"
	"fmt"

	"uzi/internal/domain"
	nodeEntity "uzi/internal/repository/node/entity"
	segmentEntity "uzi/internal/repository/segment/entity"

	"github.com/google/uuid"
)

func (s *service) GetNodesWithSegmentsByImageID(ctx context.Context, id uuid.UUID) ([]domain.Node, []domain.Segment, error) {
	segments, err := s.dao.NewSegmentQuery(ctx).GetSegmentsByImageID(id)
	if err != nil {
		return nil, nil, fmt.Errorf("get segments by image_id: %w", err)
	}

	nodes, err := s.dao.NewNodeQuery(ctx).GetNodesByImageID(id)
	if err != nil {
		return nil, nil, fmt.Errorf("get nodes by image_id: %w", err)
	}

	return nodeEntity.Node{}.SliceToDomain(nodes), segmentEntity.Segment{}.SliceToDomain(segments), nil
}
