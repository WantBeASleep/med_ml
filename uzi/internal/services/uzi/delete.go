package uzi

import (
	"context"
	"errors"

	"uzi/internal/domain"
	"uzi/internal/repository/entity"

	"github.com/google/uuid"
)

func (s *service) DeleteUzi(ctx context.Context, id uuid.UUID) error {
	err := s.dao.NewUziQuery(ctx).DeleteUzi(id)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return domain.ErrNotFound
		}
		return err
	}
	return nil
}
