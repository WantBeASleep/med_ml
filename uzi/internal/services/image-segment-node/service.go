// Доменные область слишком сильно пересекается
// в рамках image, segment, node
// проще всего вынести в отдельный пакет надстройки
package imagesegmentnode

import (
	"context"
	"fmt"

	"uzi/internal/domain"
	"uzi/internal/repository"
	"uzi/internal/repository/entity"

	"github.com/google/uuid"
)

type Service interface {
	GetImageSegmentsWithNodes(ctx context.Context, id uuid.UUID) ([]domain.Node, []domain.Segment, error)
}

type service struct {
	dao repository.DAO
}

func New(dao repository.DAO) Service {
	return &service{dao: dao}
}

func (s *service) GetImageSegmentsWithNodes(ctx context.Context, id uuid.UUID) ([]domain.Node, []domain.Segment, error) {
	segments, err := s.dao.NewSegmentQuery(ctx).GetSegmentsByImageID(id)
	if err != nil {
		return nil, nil, fmt.Errorf("get segments by image_id: %w", err)
	}

	// TODO: переделать на запросе без JOIN
	nodes, err := s.dao.NewNodeQuery(ctx).GetNodesByImageID(id)
	if err != nil {
		return nil, nil, fmt.Errorf("get nodes by image_id: %w", err)
	}

	return entity.Node{}.SliceToDomain(nodes), entity.Segment{}.SliceToDomain(segments), nil
}
