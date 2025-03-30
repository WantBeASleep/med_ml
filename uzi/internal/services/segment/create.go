package segment

import (
	"context"
	"fmt"

	"uzi/internal/domain"
	segmentEntity "uzi/internal/repository/segment/entity"

	"github.com/google/uuid"
)

func (s *service) CreateManualSegment(ctx context.Context, arg CreateSegmentArg) (uuid.UUID, error) {
	node, err := s.dao.NewNodeQuery(ctx).GetNodeByID(arg.NodeID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("get node by id: %w", err)
	}

	if node.Ai {
		return uuid.Nil, ErrAddSegmentToAiNode
	}

	segment := domain.Segment{
		Id:       uuid.New(),
		ImageID:  arg.ImageID,
		NodeID:   arg.NodeID,
		Contor:   arg.Contor,
		Ai:       false,
		Tirads23: arg.Tirads23,
		Tirads4:  arg.Tirads4,
		Tirads5:  arg.Tirads5,
	}
	if err := s.dao.NewSegmentQuery(ctx).InsertSegments(segmentEntity.Segment{}.FromDomain(segment)); err != nil {
		return uuid.Nil, err
	}

	return segment.Id, nil
}
