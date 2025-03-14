package segment

import (
	"context"

	"github.com/google/uuid"
)

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.adapters.Uzi.DeleteSegment(ctx, id)
}
