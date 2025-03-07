package image

import (
	"context"

	"uzi/internal/domain"
	"uzi/internal/repository/image/entity"

	"github.com/google/uuid"
)

func (s *service) GetImagesByUziID(ctx context.Context, id uuid.UUID) ([]domain.Image, error) {
	images, err := s.dao.NewImageQuery(ctx).GetImagesByUziID(id)
	if err != nil {
		return nil, err
	}

	return entity.Image{}.SliceToDomain(images), nil
}
