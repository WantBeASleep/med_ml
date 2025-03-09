package device

import (
	"context"
	
	domain "gateway/internal/domain/uzi"
	"gateway/internal/adapters"
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

