package doctor

import (
	"context"

	domain "composition-api/internal/domain/med"

	"github.com/google/uuid"
)

func (s *service) GetDoctor(ctx context.Context, id uuid.UUID) (domain.Doctor, error) {
	doctor, err := s.adapters.Med.GetDoctor(ctx, id)
	if err != nil {
		return domain.Doctor{}, err
	}

	return doctor, nil
}
