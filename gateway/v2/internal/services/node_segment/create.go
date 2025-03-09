package node_segment

import (
	"context"

	"github.com/google/uuid"

	adapter "gateway/internal/adapters/uzi"
)

func (s *service) CreateNodeWithSegment(ctx context.Context, arg CreateNodeWithSegmentArg) (uuid.UUID, []uuid.UUID, error) {
	nodeID, segmentIDs, err := s.adapters.Uzi.CreateNodeWithSegments(ctx, adapter.CreateNodeWithSegmentsIn{
		Node:     arg.Node,
		Segments: arg.Segments,
	})
	if err != nil {
		return uuid.Nil, nil, err
	}
	return nodeID, segmentIDs, nil
}
