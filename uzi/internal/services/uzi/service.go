package uzi

import (
	"context"

	"uzi/internal/domain"
	"uzi/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	CreateUzi(ctx context.Context, arg CreateUziArg) (uuid.UUID, error)

	GetUziByID(ctx context.Context, id uuid.UUID) (domain.Uzi, error)
	GetUzisByExternalID(ctx context.Context, externalID uuid.UUID) ([]domain.Uzi, error)
	GetUzisByAuthor(ctx context.Context, author uuid.UUID) ([]domain.Uzi, error)
	GetUziEchographicsByID(ctx context.Context, id uuid.UUID) (domain.Echographic, error)

	UpdateUzi(ctx context.Context, arg UpdateUziArg) (domain.Uzi, error)
	UpdateEchographic(ctx context.Context, arg UpdateEchographicArg) (domain.Echographic, error)

	DeleteUzi(ctx context.Context, id uuid.UUID) error
}

type service struct {
	dao repository.DAO
}

func New(
	dao repository.DAO,
) Service {
	return &service{
		dao: dao,
	}
}
