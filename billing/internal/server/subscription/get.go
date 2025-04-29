package subscription

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "billing/internal/generated/grpc/service"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *handler) IsUserHasActiveSubscription(ctx context.Context, in *pb.IsUserHasActiveSubscriptionIn) (*pb.IsUserHasActiveSubscriptionOut, error) {
	userID, err := uuid.Parse(in.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "user_id is not a valid uuid: %s", err.Error())
	}

	hasActive, err := h.services.Subscription.IsUserHasActiveSubscrption(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check active subscription: %s", err.Error())
	}

	return &pb.IsUserHasActiveSubscriptionOut{HasActiveSubscription: hasActive}, nil
}

func (h *handler) GetUserActiveSubscription(ctx context.Context, in *pb.GetUserActiveSubscriptionIn) (*pb.GetUserActiveSubscriptionOut, error) {
	userID, err := uuid.Parse(in.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "user_id is not a valid uuid: %s", err.Error())
	}

	subscription, err := h.services.Subscription.GetUserActiveSubscrption(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get active subscription: %s", err.Error())
	}

	if subscription == nil {
		return nil, status.Errorf(codes.NotFound, "no active subscription found")
	}

	subscriptionProto := &pb.Subscription{
		SubscriptionId: subscription.ID.String(),
		TariffPlanId:   subscription.TariffPlanID.String(),
		Status:         pb.SubscriptionStatus(pb.SubscriptionStatus_value[string(subscription.Status)]),
		StartDate:      timestamppb.New(subscription.StartDate),
		EndDate:        timestamppb.New(subscription.EndDate),
	}

	return &pb.GetUserActiveSubscriptionOut{
		Subscription: subscriptionProto,
	}, nil
}
