package subscription

import (
	"context"

	pb "billing/internal/generated/grpc/service"
	"billing/internal/services"
)

type SubscriptionHandler interface {
	PurchaseSubscription(ctx context.Context, req *pb.PurchaseSubscriptionIn) (*pb.PurchaseSubscriptionOut, error)
	// SetSubscriptionStatus(ctx context.Context, req *pb.SetSubscriptionStatusIn) (*pb.Empty, error)
	IsUserHasActiveSubscription(ctx context.Context, req *pb.IsUserHasActiveSubscriptionIn) (*pb.IsUserHasActiveSubscriptionOut, error)
	GetUserActiveSubscription(ctx context.Context, req *pb.GetUserActiveSubscriptionIn) (*pb.GetUserActiveSubscriptionOut, error)
}

type handler struct {
	services *services.Services
}

func New(
	services *services.Services,
) SubscriptionHandler {
	return &handler{
		services: services,
	}
}
