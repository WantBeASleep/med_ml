package node_segment

import (
	"context"

	"github.com/google/uuid"

	"gateway/internal/adapters"
	domain "gateway/internal/domain/uzi"
)

type Service interface {
	CreateNodeWithSegment(ctx context.Context, arg CreateNodeWithSegmentArg) (uuid.UUID, []uuid.UUID, error)
	GetNodeWithSegmentsByImageID(ctx context.Context, imageID uuid.UUID) ([]domain.Node, []domain.Segment, error)
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
