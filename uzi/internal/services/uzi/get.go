package uzi

import (
	"context"

	"uzi/internal/domain"
	uziEntity "uzi/internal/repository/uzi/entity"

	"github.com/google/uuid"
)

func (s *service) GetUziByID(ctx context.Context, id uuid.UUID) (domain.Uzi, error) {
	uzi, err := s.dao.NewUziQuery(ctx).GetUziByID(id)
	if err != nil {
		return domain.Uzi{}, err
	}

	return uzi.ToDomain(), nil
}

func (s *service) GetUzisByExternalID(ctx context.Context, externalID uuid.UUID) ([]domain.Uzi, error) {
	uzis, err := s.dao.NewUziQuery(ctx).GetUzisByExternalID(externalID)
	if err != nil {
		return nil, err
	}

	return uziEntity.Uzi{}.SliceToDomain(uzis), nil
}

func (s *service) GetUziEchographicsByID(ctx context.Context, id uuid.UUID) (domain.Echographic, error) {
	echographics, err := s.dao.NewEchographicQuery(ctx).GetEchographicByID(id)
	if err != nil {
		return domain.Echographic{}, err
	}

	return echographics.ToDomain(), nil
}
