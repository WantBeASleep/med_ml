package image

import (
	"context"

	"github.com/google/uuid"

	domain "composition-api/internal/domain/uzi"
)

func (s *service) GetImagesByUziID(ctx context.Context, uziID uuid.UUID) ([]domain.Image, error) {
	images, err := s.adapters.Uzi.GetImagesByUziId(ctx, uziID)
	if err != nil {
		return nil, err
	}
	return images, nil
}
