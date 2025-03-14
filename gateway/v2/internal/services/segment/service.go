package segment

import (
	"context"

	"github.com/google/uuid"

	"gateway/internal/adapters"
	domain "gateway/internal/domain/uzi"
)

type Service interface {
	Create(ctx context.Context, arg CreateSegmentArg) (uuid.UUID, error)
	GetByNodeID(ctx context.Context, nodeID uuid.UUID) ([]domain.Segment, error)
	Update(ctx context.Context, arg UpdateSegmentArg) (domain.Segment, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type service struct {
	adapters *adapters.Adapters
}

func New(
	adapters *adapters.Adapters,
) Service {
	return &service{
		adapters: adapters,
	}
}
