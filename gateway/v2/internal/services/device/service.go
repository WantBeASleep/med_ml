package device

import (
	"context"

	"gateway/internal/adapters"
	domain "gateway/internal/domain/uzi"
)

type Service interface {
	Create(ctx context.Context, name string) (int, error)
	GetAll(ctx context.Context) ([]domain.Device, error)
}

type service struct {
	adapters *adapters.Adapters
}

func New(adapters *adapters.Adapters) Service {
	return &service{adapters: adapters}
}
