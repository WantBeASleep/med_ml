package subscription

import (
	"context"

	"github.com/google/uuid"

	"composition-api/internal/adapters"
	domain "composition-api/internal/domain/billing"
)

type Service interface {
	PurchaseSubscription(ctx context.Context, tariffPlanID, paymentProviderID, userID uuid.UUID) (domain.PurchaseSubscriptionResponse, error)
	IsUserHasActiveSubscription(ctx context.Context, userID uuid.UUID) (bool, error)
	GetUserActiveSubscription(ctx context.Context, userID uuid.UUID) (domain.Subscription, error)
}

type service struct {
	adapters *adapters.Adapters
}

func New(adapters *adapters.Adapters) Service {
	return &service{
		adapters: adapters,
	}
}
