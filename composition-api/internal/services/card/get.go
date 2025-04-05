package card

import (
	"context"

	domain "composition-api/internal/domain/med"

	"github.com/google/uuid"
)

func (s *service) GetCard(ctx context.Context, doctorID, patientID uuid.UUID) (domain.Card, error) {
	card, err := s.adapters.Med.GetCard(ctx, doctorID, patientID)
	if err != nil {
		return domain.Card{}, err
	}

	return card, nil
}
