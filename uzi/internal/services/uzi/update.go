package uzi

import (
	"context"
	"fmt"

	"uzi/internal/domain"
	echographicEntity "uzi/internal/repository/echographic/entity"
	uziEntity "uzi/internal/repository/uzi/entity"
)

func (s *service) UpdateUzi(ctx context.Context, arg UpdateUziArg) (domain.Uzi, error) {
	uzi, err := s.GetUziByID(ctx, arg.Id)
	if err != nil {
		return domain.Uzi{}, fmt.Errorf("get uzi by id: %w", err)
	}
	arg.UpdateDomain(&uzi)

	if err := s.dao.NewUziQuery(ctx).UpdateUzi(uziEntity.Uzi{}.FromDomain(uzi)); err != nil {
		return domain.Uzi{}, fmt.Errorf("update uzi: %w", err)
	}

	return uzi, nil
}

func (s *service) UpdateEchographic(ctx context.Context, arg UpdateEchographicArg) (domain.Echographic, error) {
	echographic, err := s.GetUziEchographicsByID(ctx, arg.Id)
	if err != nil {
		return domain.Echographic{}, fmt.Errorf("get uzi by id: %w", err)
	}
	arg.UpdateDomain(&echographic)

	if err := s.dao.NewEchographicQuery(ctx).UpdateEchographic(echographicEntity.Echographic{}.FromDomain(echographic)); err != nil {
		return domain.Echographic{}, fmt.Errorf("update echographic: %w", err)
	}

	return echographic, nil
}
