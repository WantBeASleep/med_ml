package billing

import (
	"context"
	"fmt"

	adapter_errors "composition-api/internal/adapters/errors"
	billing_domain "composition-api/internal/domain/billing"

	"composition-api/internal/adapters/billing/mappers"
	pb "composition-api/internal/generated/grpc/clients/billing"

	"github.com/google/uuid"
)

func (a *adapter) PurchaseSubscription(ctx context.Context, tariffPlanID, paymentProviderID, userID uuid.UUID) (billing_domain.PurchaseSubscriptionResponse, error) {
	res, err := a.client.PurchaseSubscription(ctx, &pb.PurchaseSubscriptionIn{
		TariffPlanId:      tariffPlanID.String(),
		PaymentProviderId: paymentProviderID.String(),
		UserId:            userID.String(),
	})
	if err != nil {
		return billing_domain.PurchaseSubscriptionResponse{}, fmt.Errorf("failed to purchase subscription: %w", err)
	}
	return mappers.PurchaseSubscription{}.Domain(res), nil
}

func (a *adapter) IsUserHasActiveSubscription(ctx context.Context, userID uuid.UUID) (bool, error) {
	res, err := a.client.IsUserHasActiveSubscription(ctx, &pb.IsUserHasActiveSubscriptionIn{UserId: userID.String()})
	if err != nil {
		return false, fmt.Errorf("failed to check active subscription: %w", err)
	}
	return res.HasActiveSubscription, nil
}

func (a *adapter) GetUserActiveSubscription(ctx context.Context, userID uuid.UUID) (billing_domain.Subscription, error) {
	res, err := a.client.GetUserActiveSubscription(ctx, &pb.GetUserActiveSubscriptionIn{UserId: userID.String()})
	if err != nil {
		return billing_domain.Subscription{}, adapter_errors.HandleGRPCError(err)
	}
	return mappers.Subscription{}.Domain(res.Subscription), nil
}
