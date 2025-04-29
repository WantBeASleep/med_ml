package subscription

import (
	"context"

	"github.com/google/uuid"

	domain "composition-api/internal/domain/billing"
)

func (s *service) PurchaseSubscription(ctx context.Context, tariffPlanID, paymentProviderID, userID uuid.UUID) (domain.PurchaseSubscriptionResponse, error) {
	return s.adapters.Billing.PurchaseSubscription(ctx, tariffPlanID, paymentProviderID, userID)
}
