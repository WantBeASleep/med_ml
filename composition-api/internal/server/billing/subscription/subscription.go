package subscription

import (
	"context"
	"errors"

	adapter_errors "composition-api/internal/adapters/errors"
	"composition-api/internal/server/security"

	api "composition-api/internal/generated/http/api"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
)

func (h *handler) SubscriptionsPurchasePost(ctx context.Context, req *api.PurchaseSubscriptionRequest) (api.SubscriptionsPurchasePostRes, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	purchaseSubscriptionResponse, err := h.services.SubscriptionService.PurchaseSubscription(ctx, req.TariffPlanID, req.PaymentProviderID, userID)
	if err != nil {
		return nil, err
	}

	return pointer.To(api.PurchaseSubscriptionResponse{
		SubscriptionID:  purchaseSubscriptionResponse.SubscriptionID,
		ConfirmationURL: purchaseSubscriptionResponse.ConfirmationURL,
	}), nil
}

func (h *handler) SubscriptionsCheckActiveGet(ctx context.Context) (api.SubscriptionsCheckActiveGetRes, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	hasActive, err := h.services.SubscriptionService.IsUserHasActiveSubscription(ctx, userID)
	if err != nil {
		return nil, err
	}

	return pointer.To(api.SubscriptionsCheckActiveGetOK{
		HasActiveSubscription: api.NewOptBool(hasActive),
	}), nil
}

func (h *handler) SubscriptionsGetActiveGet(ctx context.Context) (api.SubscriptionsGetActiveGetRes, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		if errors.Is(err, adapter_errors.ErrNotFound) {
			return pointer.To(
				api.SubscriptionsGetActiveGetNotFound(
					api.ErrorStatusCode{
						StatusCode: 404,
						Response: api.Error{
							Code:    404,
							Message: err.Error(),
						},
					},
				),
			), nil
		}
	}
	subscription, err := h.services.SubscriptionService.GetUserActiveSubscription(ctx, userID)
	if err != nil {
		return nil, err
	}

	return pointer.To(api.Subscription{
		ID:           subscription.ID,
		TariffPlanID: subscription.TariffPlanID,
		Status:       api.SubscriptionStatus(subscription.Status),
		StartDate:    subscription.StartDate,
		EndDate:      subscription.EndDate,
	}), nil
}

func getUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	token, err := security.ParseToken(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	return token.Id, nil
}
