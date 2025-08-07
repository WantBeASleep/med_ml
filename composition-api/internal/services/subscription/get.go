package subscription

import (
	"context"

	"github.com/google/uuid"

	domain "composition-api/internal/domain/billing"
)

func (s *service) IsUserHasActiveSubscription(ctx context.Context, userID uuid.UUID) (bool, error) {
	return s.adapters.Billing.IsUserHasActiveSubscription(ctx, userID)
}

func (s *service) GetUserActiveSubscription(ctx context.Context, userID uuid.UUID) (domain.Subscription, error) {
	return s.adapters.Billing.GetUserActiveSubscription(ctx, userID)
}
