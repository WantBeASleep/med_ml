package card

import (
	"context"
	"fmt"

	"med/internal/domain"

	"github.com/google/uuid"
)

func (s *service) GetCard(ctx context.Context, doctorID, patientID uuid.UUID) (domain.Card, error) {
	card, err := s.dao.NewCardQuery(ctx).GetCardByPK(doctorID, patientID)
	if err != nil {
		return domain.Card{}, fmt.Errorf("get card by id: %w", err)
	}

	return card.ToDomain(), nil
}
