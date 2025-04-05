package segment

import (
	"context"
	"errors"

	"uzi/internal/domain"
	"uzi/internal/repository"

	"github.com/google/uuid"
)

var (
	ErrChangeAiSegment    = errors.New("change ai segment not allowed")
	ErrAddSegmentToAiNode = errors.New("add segment to ai node not allowed")
)

type Service interface {
	CreateManualSegment(ctx context.Context, arg CreateSegmentArg) (uuid.UUID, error)

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
