package card

import (
	"context"

	domain "composition-api/internal/domain/med"
)

func (s *service) CreateCard(ctx context.Context, card domain.Card) (domain.Card, error) {
	return s.adapters.Med.CreateCard(ctx, card)
}
