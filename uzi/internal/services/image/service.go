package image

import (
	"context"

	"uzi/internal/adapters/dbus"

	"uzi/internal/domain"
	"uzi/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	SplitUzi(ctx context.Context, id uuid.UUID) error

	GetImagesByUziID(ctx context.Context, id uuid.UUID) ([]domain.Image, error)
}

type service struct {
	dao     repository.DAO
	adapter dbus.DbusAdapter
}

func New(
	dao repository.DAO,
	adapter dbus.DbusAdapter,
) Service {
	return &service{
		dao:     dao,
		adapter: adapter,
	}
}
