package card

import (
	"context"
	"errors"
	"fmt"

	"med/internal/domain"
	centity "med/internal/repository/card/entity"
	"med/internal/repository/entity"

	"github.com/google/uuid"
)

func (s *service) UpdateCard(ctx context.Context, doctorID, patientID uuid.UUID, update UpdateCardArg) (domain.Card, error) {
	cardQuery := s.dao.NewCardQuery(ctx)

	card, err := s.GetCard(ctx, doctorID, patientID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return domain.Card{}, domain.ErrNotFound
		}
		return domain.Card{}, fmt.Errorf("get card: %w", err)
	}
	update.Update(&card)

	if err := cardQuery.UpdateCard(centity.Card{}.FromDomain(card)); err != nil {
		switch {
		case errors.Is(err, entity.ErrConflict):
			return domain.Card{}, domain.ErrConflict
		case errors.Is(err, entity.ErrValidation):
			return domain.Card{}, domain.ErrUnprocessableEntity
		default:
			return domain.Card{}, fmt.Errorf("update card: %w", err)
		}
	}

	return card, nil
}
