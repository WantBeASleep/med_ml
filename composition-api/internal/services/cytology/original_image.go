package cytology

import (
	"context"

	"github.com/google/uuid"

	"composition-api/internal/adapters/cytology"
	domain "composition-api/internal/domain/cytology"
)

func (s *service) CreateOriginalImage(ctx context.Context, arg CreateOriginalImageArg) (uuid.UUID, error) {
	return s.adapters.Cytology.CreateOriginalImage(ctx, cytology.CreateOriginalImageIn{
		CytologyID:  arg.CytologyID,
		File:        arg.File,
		ContentType: arg.ContentType,
		DelayTime:   arg.DelayTime,
	})
}

func (s *service) GetOriginalImageById(ctx context.Context, id uuid.UUID) (domain.OriginalImage, error) {
	return s.adapters.Cytology.GetOriginalImageById(ctx, id)
}

func (s *service) GetOriginalImagesByCytologyId(ctx context.Context, id uuid.UUID) ([]domain.OriginalImage, error) {
	return s.adapters.Cytology.GetOriginalImagesByCytologyId(ctx, id)
}

func (s *service) UpdateOriginalImage(ctx context.Context, arg UpdateOriginalImageArg) (domain.OriginalImage, error) {
	return s.adapters.Cytology.UpdateOriginalImage(ctx, cytology.UpdateOriginalImageIn{
		Id:         arg.Id,
		DelayTime:  arg.DelayTime,
		ViewedFlag: arg.ViewedFlag,
	})
}
