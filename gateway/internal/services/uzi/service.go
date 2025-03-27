package uzi

import (
	"context"

	"github.com/google/uuid"

	"composition-api/internal/adapters"
	dbus "composition-api/internal/dbus/producers"
	domain "composition-api/internal/domain/uzi"
	"composition-api/internal/repository"
)

type Service interface {
	Create(ctx context.Context, arg CreateUziArg) (uuid.UUID, error)

	GetByID(ctx context.Context, id uuid.UUID) (domain.Uzi, error)
	GetByExternalID(ctx context.Context, externalID uuid.UUID) ([]domain.Uzi, error)
	GetEchographicsByID(ctx context.Context, id uuid.UUID) (domain.Echographic, error)

	Update(ctx context.Context, arg UpdateUziArg) (domain.Uzi, error)
	UpdateEchographics(ctx context.Context, arg domain.Echographic) (domain.Echographic, error)
}

type service struct {
	adapters *adapters.Adapters
	dao      repository.DAO
	dbus     dbus.Producer
}

func New(
	adapters *adapters.Adapters,
	dao repository.DAO,
	dbus dbus.Producer,
) Service {
	return &service{
		adapters: adapters,
		dao:      dao,
		dbus:     dbus,
	}
}
