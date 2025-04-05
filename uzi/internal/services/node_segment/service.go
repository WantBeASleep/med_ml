// Доменные области слишком сильно пересекаются в рамках node, segment
// проще всего вынести в отдельный пакет надстройки
package node_segment

import (
	"context"

	"uzi/internal/domain"
	"uzi/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	SaveProcessedNodesWithSegments(ctx context.Context, uziID uuid.UUID, arg []CreateNodesWithSegmentsArg) error
	CreateManualNodesWithSegments(ctx context.Context, uziID uuid.UUID, arg []CreateNodesWithSegmentsArg) ([]CreateNodesWithSegmentsID, error)

	GetNodesWithSegmentsByImageID(ctx context.Context, id uuid.UUID) ([]domain.Node, []domain.Segment, error)

	DeleteNode(ctx context.Context, id uuid.UUID) error
	DeleteSegment(ctx context.Context, id uuid.UUID) error
}

type service struct {
	dao repository.DAO
}

func New(dao repository.DAO) Service {
	return &service{dao: dao}
}
