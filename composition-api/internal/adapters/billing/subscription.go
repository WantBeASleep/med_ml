package billing

import (
	"context"
	"fmt"

	adapter_errors "composition-api/internal/adapters/errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"composition-api/internal/adapters/billing/mappers"
	domain "composition-api/internal/domain/billing"
	pb "composition-api/internal/generated/grpc/clients/billing"

	"github.com/google/uuid"
)

func (a *adapter) PurchaseSubscription(ctx context.Context, tariffPlanID, paymentProviderID, userID uuid.UUID) (domain.PurchaseSubscriptionResponse, error) {
	res, err := a.client.PurchaseSubscription(ctx, &pb.PurchaseSubscriptionIn{
		TariffPlanId:      tariffPlanID.String(),
		PaymentProviderId: paymentProviderID.String(),
		UserId:            userID.String(),
	})
	if err != nil {
		return domain.PurchaseSubscriptionResponse{}, fmt.Errorf("failed to purchase subscription: %w", err)
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

func (a *adapter) GetUserActiveSubscription(ctx context.Context, userID uuid.UUID) (domain.Subscription, error) {
	res, err := a.client.GetUserActiveSubscription(ctx, &pb.GetUserActiveSubscriptionIn{UserId: userID.String()})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return domain.Subscription{}, fmt.Errorf("unknown error: %w", err)
		}
		switch st.Code() {
		case codes.NotFound:
			return domain.Subscription{}, adapter_errors.ErrNotFound
		default:
			return domain.Subscription{}, err
		}
	}
	return mappers.Subscription{}.Domain(res.Subscription), nil
}
