package uzi

import (
	"context"
	"errors"
	"fmt"
	"time"

	"uzi/internal/domain"
	echographicEntity "uzi/internal/repository/echographic/entity"
	"uzi/internal/repository/entity"
	uziEntity "uzi/internal/repository/uzi/entity"

	"github.com/google/uuid"
)

func (s *service) CreateUzi(ctx context.Context, arg CreateUziArg) (uuid.UUID, error) {
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("begin transaction: %w", err)
	}
	defer func() { _ = s.dao.RollbackTx(ctx) }()

	uzi := domain.Uzi{
		Id:          uuid.New(),
		Projection:  arg.Projection,
		Checked:     false,
		ExternalID:  arg.ExternalID,
		Author:      arg.Author,
		DeviceID:    arg.DeviceID,
		Status:      domain.UziStatusNew,
		Description: arg.Description,
		CreateAt:    time.Now(),
	}

	if err := s.dao.NewUziQuery(ctx).InsertUzi(uziEntity.Uzi{}.FromDomain(uzi)); err != nil {
		var valErr *entity.DBValidationError
		if errors.As(err, &valErr) {
			return uuid.Nil, domain.ErrUnprocessableEntity
		}
		return uuid.Nil, fmt.Errorf("insert uzi: %w", err)
	}

	if err := s.dao.NewEchographicQuery(ctx).InsertEchographic(echographicEntity.Echographic{Id: uzi.Id}); err != nil {
		var valErr *entity.DBValidationError
		if errors.As(err, &valErr) {
			return uuid.Nil, domain.ErrUnprocessableEntity
		}
		return uuid.Nil, fmt.Errorf("insert echographic: %w", err)
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return uuid.Nil, fmt.Errorf("commit transaction: %w", err)
	}

	return uzi.Id, nil
}
