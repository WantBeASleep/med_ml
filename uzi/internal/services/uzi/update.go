package uzi

import (
	"context"
	"errors"
	"fmt"

	"uzi/internal/domain"
	echographicEntity "uzi/internal/repository/echographic/entity"
	"uzi/internal/repository/entity"
	uziEntity "uzi/internal/repository/uzi/entity"
)

func (s *service) UpdateUzi(ctx context.Context, arg UpdateUziArg) (domain.Uzi, error) {
	uzi, err := s.GetUziByID(ctx, arg.Id)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return domain.Uzi{}, domain.ErrNotFound
		}
		return domain.Uzi{}, fmt.Errorf("get uzi by id: %w", err)
	}
	arg.UpdateDomain(&uzi)

	if err := s.dao.NewUziQuery(ctx).UpdateUzi(uziEntity.Uzi{}.FromDomain(uzi)); err != nil {
		var dbErr *entity.DBConflictError
		if errors.As(err, &dbErr) {
			return domain.Uzi{}, domain.ErrConflict
		}
		var valErr *entity.DBValidationError
		if errors.As(err, &valErr) {
			return domain.Uzi{}, domain.ErrUnprocessableEntity
		}
		return domain.Uzi{}, fmt.Errorf("update uzi: %w", err)
	}

	return uzi, nil
}

func (s *service) UpdateEchographic(ctx context.Context, arg UpdateEchographicArg) (domain.Echographic, error) {
	echographic, err := s.GetUziEchographicsByID(ctx, arg.Id)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return domain.Echographic{}, domain.ErrNotFound
		}
		return domain.Echographic{}, fmt.Errorf("get uzi by id: %w", err)
	}
	arg.UpdateDomain(&echographic)

	if err := s.dao.NewEchographicQuery(ctx).UpdateEchographic(echographicEntity.Echographic{}.FromDomain(echographic)); err != nil {
		var dbErr *entity.DBConflictError
		if errors.As(err, &dbErr) {
			return domain.Echographic{}, domain.ErrConflict
		}
		var valErr *entity.DBValidationError
		if errors.As(err, &valErr) {
			return domain.Echographic{}, domain.ErrUnprocessableEntity
		}
		return domain.Echographic{}, fmt.Errorf("update echographic: %w", err)
	}

	return echographic, nil
}
