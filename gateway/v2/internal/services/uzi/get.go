package uzi

import (
	"context"

	"github.com/google/uuid"

	domain "gateway/internal/domain/uzi"
)

func (s *service) GetByID(ctx context.Context, id uuid.UUID) (domain.Uzi, error) {
	uzi, err := s.adapters.Uzi.GetUziById(ctx, id)
	if err != nil {
		return domain.Uzi{}, err
	}
	return uzi, nil
}

func (s *service) GetByExternalID(ctx context.Context, externalID uuid.UUID) ([]domain.Uzi, error) {
	uzis, err := s.adapters.Uzi.GetUzisByExternalId(ctx, externalID)
	if err != nil {
		return nil, err
	}
	return uzis, nil
}

func (s *service) GetEchographicsByID(ctx context.Context, id uuid.UUID) (domain.Echographic, error) {
	echographics, err := s.adapters.Uzi.GetEchographicByUziId(ctx, id)
	if err != nil {
		return domain.Echographic{}, err
	}
	return echographics, nil
}
