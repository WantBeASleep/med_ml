package image

import (
	"context"

	"github.com/google/uuid"

	"gateway/internal/adapters"
	domain "gateway/internal/domain/uzi"
)

type Service interface {
	GetImagesByUziID(ctx context.Context, uziID uuid.UUID) ([]domain.Image, error)
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
