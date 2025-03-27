package segment

import (
	"context"

	"github.com/google/uuid"

	domain "composition-api/internal/domain/uzi"
)

func (s *service) GetByNodeID(ctx context.Context, nodeID uuid.UUID) ([]domain.Segment, error) {
	return s.adapters.Uzi.GetSegmentsByNodeId(ctx, nodeID)
}
