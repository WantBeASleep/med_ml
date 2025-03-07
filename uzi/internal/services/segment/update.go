package segment

import (
	"context"
	"fmt"

	"uzi/internal/domain"
	segmentEntity "uzi/internal/repository/segment/entity"
)

func (s *service) UpdateSegment(ctx context.Context, arg UpdateSegmentArg) (domain.Segment, error) {
	segmentQuery := s.dao.NewSegmentQuery(ctx)
	segmentDB, err := segmentQuery.GetSegmentByID(arg.Id)
	if err != nil {
		return domain.Segment{}, fmt.Errorf("get segment by id: %w", err)
	}
	segment := segmentDB.ToDomain()
	arg.UpdateDomain(&segment)

	if err := segmentQuery.UpdateSegment(segmentEntity.Segment{}.FromDomain(segment)); err != nil {
		return domain.Segment{}, fmt.Errorf("update segment: %w", err)
	}

	return segment, nil
}
