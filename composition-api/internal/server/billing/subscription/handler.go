package subscription

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type SubscriptionHandler interface {
	SubscriptionsPurchasePost(ctx context.Context, req *api.PurchaseSubscriptionRequest) (api.SubscriptionsPurchasePostRes, error)
	SubscriptionsCheckActiveGet(ctx context.Context) (api.SubscriptionsCheckActiveGetRes, error)
	SubscriptionsGetActiveGet(ctx context.Context) (api.SubscriptionsGetActiveGetRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) SubscriptionHandler {
	return &handler{
		services: services,
	}
}
