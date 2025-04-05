package segment

import (
	"context"

	"github.com/google/uuid"

	adapter "composition-api/internal/adapters/uzi"
)

func (s *service) Create(ctx context.Context, arg CreateSegmentArg) (uuid.UUID, error) {
	id, err := s.adapters.Uzi.CreateSegment(ctx, adapter.CreateSegmentIn{
		ImageID:   arg.ImageID,
		NodeID:    arg.NodeID,
		Contor:    arg.Contor,
		Tirads_23: arg.Tirads_23,
		Tirads_4:  arg.Tirads_4,
		Tirads_5:  arg.Tirads_5,
	})
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
