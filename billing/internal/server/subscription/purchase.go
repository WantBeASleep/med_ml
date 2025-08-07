package subscription

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "billing/internal/generated/grpc/service"
)

func (h *handler) PurchaseSubscription(ctx context.Context, in *pb.PurchaseSubscriptionIn) (*pb.PurchaseSubscriptionOut, error) {
	tariffPlanID, err := uuid.Parse(in.TariffPlanId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "tariff_plan_id is not a valid uuid: %s", err.Error())
	}

	paymentProviderID, err := uuid.Parse(in.PaymentProviderId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "payment_provider_id is not a valid uuid: %s", err.Error())
	}

	userID, err := uuid.Parse(in.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "user_id is not a valid uuid: %s", err.Error())
	}

	subscriptionID, confirmationURL, err := h.services.Subscription.PurchaseSubscrption(ctx, tariffPlanID, paymentProviderID, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to purchase subscription: %s", err.Error())
	}

	return &pb.PurchaseSubscriptionOut{
		SubscriptionId:  subscriptionID.String(),
		ConfirmationUrl: confirmationURL,
	}, nil
}
