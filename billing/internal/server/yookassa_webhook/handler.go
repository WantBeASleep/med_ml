package yookassa_webhook

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "billing/internal/generated/grpc/service"
	"billing/internal/services"
	"billing/internal/services/notification_manager"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type YookassaWebhookHandler interface {
	HandleYookassaWebhook(ctx context.Context, req *pb.YookassaWebhookRequest) (*emptypb.Empty, error)
}

type handler struct {
	services *services.Services
}

func New(
	services *services.Services,
) YookassaWebhookHandler {
	return &handler{
		services: services,
	}
}

func (h *handler) HandleYookassaWebhook(ctx context.Context, req *pb.YookassaWebhookRequest) (*emptypb.Empty, error) {
	notification := notification_manager.UniversalNotification{
		ProviderPaymentID: req.PaymentId,
		Event:             req.Event,
		ProviderName:      req.ProviderName,
		ReceivedAt:        req.ReceivedAt.AsTime(),
		ExtraInfo:         make(map[string]interface{}),
	}

	for key, value := range req.ExtraInfo {
		notification.ExtraInfo[key] = value
	}

	switch req.Status {
	case pb.PaymentStatus_completed:
		err := h.services.NotificationManager.HandlSuccessfulNotification(ctx, notification_manager.SuccessfulNotification{UniversalNotification: notification})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to handle successful notification: %v", err)
		}
	case pb.PaymentStatus_pay_cancelled:
		err := h.services.NotificationManager.HandlCanceledNotification(ctx, notification_manager.CanceledNotification{
			UniversalNotification: notification,
			Reason:                req.Reason,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to handle canceled notification: %v", err)
		}
	case pb.PaymentStatus_waiting_for_capture:
		err := h.services.NotificationManager.HandlWaitingForCaptureNotification(ctx, notification_manager.WaitingForCaptureNotification{UniversalNotification: notification})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to handle waiting for capture notification: %v", err)
		}
	default:
		return nil, status.Errorf(codes.InvalidArgument, "Unknown status: %v", req.Status)
	}

	return &emptypb.Empty{}, nil
}
