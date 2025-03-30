package card

import (
	"context"
	"fmt"

	"med/internal/domain"
	centity "med/internal/repository/card/entity"

	"github.com/google/uuid"
)

func (s *service) UpdateCard(ctx context.Context, doctorID, patientID uuid.UUID, update UpdateCardArg) (domain.Card, error) {
	cardQuery := s.dao.NewCardQuery(ctx)

	card, err := s.GetCard(ctx, doctorID, patientID)
	if err != nil {
		return domain.Card{}, fmt.Errorf("get card: %w", err)
	}
	update.Update(&card)

	if err := cardQuery.UpdateCard(centity.Card{}.FromDomain(card)); err != nil {
		return domain.Card{}, fmt.Errorf("update card: %w", err)
	}

	return card, nil
}
