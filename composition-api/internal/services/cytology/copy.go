package cytology

import (
	"context"
	"errors"

	domain "composition-api/internal/domain/cytology"

	"github.com/google/uuid"
)

func (s *service) CopyCytologyImage(ctx context.Context, id uuid.UUID) (domain.CytologyImage, error) {
	// Получаем текущее исследование
	img, _, err := s.GetCytologyImageById(ctx, id)
	if err != nil {
		return domain.CytologyImage{}, err
	}

	// Проверяем, что это последняя версия
	if !img.IsLast {
		return domain.CytologyImage{}, errors.New("can only copy last version")
	}

	// Создаем копию через адаптер
	return s.adapters.Cytology.CopyCytologyImage(ctx, id)
}
