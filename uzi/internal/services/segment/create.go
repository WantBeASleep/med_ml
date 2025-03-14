package segment

import (
	"context"

	"uzi/internal/domain"
	segmentEntity "uzi/internal/repository/segment/entity"

	"github.com/google/uuid"
)

func (s *service) CreateSegment(ctx context.Context, arg CreateSegmentArg) (uuid.UUID, error) {
	segment := domain.Segment{
		Id:       uuid.New(),
		ImageID:  arg.ImageID,
		NodeID:   arg.NodeID,
		Contor:   arg.Contor,
		Tirads23: arg.Tirads23,
		Tirads4:  arg.Tirads4,
		Tirads5:  arg.Tirads5,
	}
	if err := s.dao.NewSegmentQuery(ctx).InsertSegments(segmentEntity.Segment{}.FromDomain(segment)); err != nil {
		return uuid.Nil, err
	}

	return segment.Id, nil
}
