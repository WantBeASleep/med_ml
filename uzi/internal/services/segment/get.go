package segment

import (
	"context"

	"uzi/internal/domain"
	segmentEntity "uzi/internal/repository/segment/entity"

	"github.com/google/uuid"
)

func (s *service) GetSegmentsByNodeID(ctx context.Context, id uuid.UUID) ([]domain.Segment, error) {
	segments, err := s.dao.NewSegmentQuery(ctx).GetSegmentsByNodeID(id)
	if err != nil {
		return nil, err
	}
	return segmentEntity.Segment{}.SliceToDomain(segments), nil
}
