package cytology_image

import (
	"context"
	"errors"
	"fmt"

	"cytology/internal/domain"
	"cytology/internal/repository/entity"

	"github.com/google/uuid"
)

func (s *service) DeleteCytologyImage(ctx context.Context, id uuid.UUID) error {
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer func() { _ = s.dao.RollbackTx(ctx) }()

	exists, err := s.dao.NewCytologyImageQuery(ctx).CheckExist(id)
	if err != nil {
		return fmt.Errorf("check exist: %w", err)
	}
	if !exists {
		return domain.ErrNotFound
	}

	if err := s.dao.NewCytologyImageQuery(ctx).DeleteCytologyImage(id); err != nil {
		var valErr *entity.DBValidationError
		if errors.As(err, &valErr) {
			return domain.ErrUnprocessableEntity
		}
		return fmt.Errorf("delete cytology image: %w", err)
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}
