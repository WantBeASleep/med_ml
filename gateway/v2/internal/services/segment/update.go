package segment

import (
	"context"

	adapter "gateway/internal/adapters/uzi"
	domain "gateway/internal/domain/uzi"
)

func (s *service) Update(ctx context.Context, arg UpdateSegmentArg) (domain.Segment, error) {
	segment, err := s.adapters.Uzi.UpdateSegment(ctx, adapter.UpdateSegmentIn{
		Id:        arg.Id,
		Tirads_23: arg.Tirads_23,
		Tirads_4:  arg.Tirads_4,
		Tirads_5:  arg.Tirads_5,
	})
	if err != nil {
		return domain.Segment{}, err
	}
	return segment, nil
}
