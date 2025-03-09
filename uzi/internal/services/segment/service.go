package segment

import (
	"context"

	"uzi/internal/domain"
	"uzi/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	CreateSegment(ctx context.Context, arg CreateSegmentArg) (uuid.UUID, error)

	GetSegmentsByNodeID(ctx context.Context, id uuid.UUID) ([]domain.Segment, error)

	UpdateSegment(ctx context.Context, arg UpdateSegmentArg) (domain.Segment, error)
}

type service struct {
	dao repository.DAO
}

func New(
	dao repository.DAO,
) Service {
	return &service{
		dao: dao,
	}
}
