package cytology_image

import (
	"context"
	"fmt"

	"cytology/internal/domain"

	"github.com/google/uuid"
)

func (s *service) GetCytologyImageHistory(ctx context.Context, id uuid.UUID) ([]domain.CytologyImage, error) {
	// Получаем текущее исследование
	img, err := s.GetCytologyImageByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get cytology image: %w", err)
	}

	// Определяем parent_prev_id для поиска
	parentPrevID := img.ParentPrevID
	if parentPrevID == nil {
		parentPrevID = &id
	}

	// Получаем все исследования с тем же parent_prev_id
	entities, err := s.dao.NewCytologyImageQuery(ctx).GetCytologyImagesByParentPrevID(*parentPrevID)
	if err != nil {
		return nil, fmt.Errorf("get cytology images by parent prev id: %w", err)
	}

	result := make([]domain.CytologyImage, 0, len(entities))
	for _, entity := range entities {
		result = append(result, entity.ToDomain())
	}

	return result, nil
}
