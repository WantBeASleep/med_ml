package cytology_image

import (
	"context"
	"errors"
	"fmt"

	"cytology/internal/domain"
	cytologyImageEntity "cytology/internal/repository/cytology_image/entity"
	"cytology/internal/repository/entity"
)

func (s *service) UpdateCytologyImage(ctx context.Context, arg UpdateCytologyImageArg) (domain.CytologyImage, error) {
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return domain.CytologyImage{}, fmt.Errorf("begin transaction: %w", err)
	}
	defer func() { _ = s.dao.RollbackTx(ctx) }()

	img, err := s.dao.NewCytologyImageQuery(ctx).GetCytologyImageByID(arg.Id)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return domain.CytologyImage{}, domain.ErrNotFound
		}
		return domain.CytologyImage{}, fmt.Errorf("get cytology image: %w", err)
	}

	domainImg := img.ToDomain()
	arg.UpdateDomain(&domainImg)

	if err := s.dao.NewCytologyImageQuery(ctx).UpdateCytologyImage(cytologyImageEntity.CytologyImage{}.FromDomain(domainImg)); err != nil {
		var valErr *entity.DBValidationError
		if errors.As(err, &valErr) {
			return domain.CytologyImage{}, domain.ErrUnprocessableEntity
		}
		return domain.CytologyImage{}, fmt.Errorf("update cytology image: %w", err)
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return domain.CytologyImage{}, fmt.Errorf("commit transaction: %w", err)
	}

	return domainImg, nil
}
