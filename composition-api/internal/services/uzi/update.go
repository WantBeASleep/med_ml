package uzi

import (
	"context"

	adapter "composition-api/internal/adapters/uzi"
	domain "composition-api/internal/domain/uzi"
)

func (s *service) Update(ctx context.Context, arg UpdateUziArg) (domain.Uzi, error) {
	uzi, err := s.adapters.Uzi.UpdateUzi(ctx, adapter.UpdateUziIn{
		Id:         arg.Id,
		Projection: arg.Projection,
		Checked:    arg.Checked,
	})
	if err != nil {
		return domain.Uzi{}, err
	}
	return uzi, nil
}

func (s *service) UpdateEchographics(ctx context.Context, arg domain.Echographic) (domain.Echographic, error) {
	echographics, err := s.adapters.Uzi.UpdateEchographic(ctx, arg)
	if err != nil {
		return domain.Echographic{}, err
	}
	return echographics, nil
}
