package payment_provider

import (
	"context"

	pb "billing/internal/generated/grpc/service"
	"billing/internal/services"

	"github.com/golang/protobuf/ptypes/empty"
)

type PaymentProviderHandler interface {
	ListPaymentProviders(ctx context.Context, _ *empty.Empty) (*pb.ListPaymentProvidersOut, error)
}

type handler struct {
	services *services.Services
}

func New(
	services *services.Services,
) PaymentProviderHandler {
	return &handler{
		services: services,
	}
}
