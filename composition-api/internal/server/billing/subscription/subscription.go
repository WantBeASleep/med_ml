package subscription

import (
	"composition-api/internal/domain"
	"composition-api/internal/server/security"
	"context"
	"errors"
	"log"
	"net/http"

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
		return nil, err
	}
	subscription, err := h.services.SubscriptionService.GetUserActiveSubscription(ctx, userID)
	if err != nil {
		log.Printf("Error retrieving subscription: %v", err)
		if errors.Is(err, domain.ErrNotFound) {
			return pointer.To(
				api.SubscriptionsGetActiveGetNotFound(
					api.ErrorStatusCode{
						StatusCode: http.StatusNotFound,
						Response: api.Error{
							Message: err.Error(),
						},
					},
				),
			), nil
		}
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
