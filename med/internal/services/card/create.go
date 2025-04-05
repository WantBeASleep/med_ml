package card

import (
	"context"
	"fmt"

	"med/internal/domain"
	centity "med/internal/repository/card/entity"
)

func (s *service) CreateCard(ctx context.Context, card domain.Card) error {
	if err := s.dao.NewCardQuery(ctx).InsertCard(centity.Card{}.FromDomain(card)); err != nil {
		return fmt.Errorf("insert card: %w", err)
	}

	return nil
}
