package image

import (
	"context"

	dbus "uzi/internal/dbus/producers"

	"uzi/internal/domain"
	"uzi/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	SplitUzi(ctx context.Context, id uuid.UUID) error

	GetImagesByUziID(ctx context.Context, id uuid.UUID) ([]domain.Image, error)
}

type service struct {
	dao  repository.DAO
	dbus dbus.Producer
}

func New(
	dao repository.DAO,
	dbus dbus.Producer,
) Service {
	return &service{
		dao:  dao,
		dbus: dbus,
	}
}
