package cytology

import (
	"context"

	domain "composition-api/internal/domain/cytology"

	"github.com/google/uuid"
)

func (s *service) GetCytologyImageHistory(ctx context.Context, id uuid.UUID) ([]domain.CytologyImage, error) {
	return s.adapters.Cytology.GetCytologyImageHistory(ctx, id)
}
