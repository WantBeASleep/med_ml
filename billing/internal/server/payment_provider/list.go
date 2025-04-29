package payment_provider

import (
	"context"

	pb "billing/internal/generated/grpc/service"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) ListPaymentProviders(ctx context.Context, _ *empty.Empty) (*pb.ListPaymentProvidersOut, error) {
	paymentProviders, err := h.services.PaymentProvider.ListPaymentProviders(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list payment providers: %s", err.Error())
	}

	var pbPaymentProviders []*pb.PaymentProvider
	for _, paymentProvider := range paymentProviders {
		pbPaymentProviders = append(pbPaymentProviders, &pb.PaymentProvider{
			Id:       paymentProvider.ID.String(),
			Name:     paymentProvider.Name,
			IsActive: paymentProvider.IsActive,
		})
	}
	return &pb.ListPaymentProvidersOut{PaymentProviders: pbPaymentProviders}, nil
}
