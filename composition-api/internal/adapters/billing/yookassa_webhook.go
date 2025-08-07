package billing

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	domain "composition-api/internal/domain/billing"
	pb "composition-api/internal/generated/grpc/clients/billing"
)

func mapStringToPaymentStatus(status string) (pb.PaymentStatus, error) {
	switch status {
	case "pending":
		return pb.PaymentStatus_pending, nil
	case "waiting_for_capture":
		return pb.PaymentStatus_waiting_for_capture, nil
	case "waiting_for_cancel":
		return pb.PaymentStatus_waiting_for_cancel, nil
	case "completed":
		return pb.PaymentStatus_completed, nil
	case "pay_cancelled":
		return pb.PaymentStatus_pay_cancelled, nil
	default:
		return 0, fmt.Errorf("unknown status: %s", status)
	}
}

func (a *adapter) HandleYookassaWebhook(ctx context.Context, req domain.YookassaWebhookRequest) error {
	status, err := mapStringToPaymentStatus(req.Status)
	if err != nil {
		return fmt.Errorf("failed to map status: %w", err)
	}

	_, err = a.client.HandleYookassaWebhook(ctx, &pb.YookassaWebhookRequest{
		Event:        req.Event,
		PaymentId:    req.PaymentID,
		Status:       status,
		ProviderName: req.ProviderName,
		ReceivedAt:   timestamppb.New(req.ReceivedAt),
		ExtraInfo:    req.ExtraInfo,
		Reason:       req.Reason,
	})
	if err != nil {
		return fmt.Errorf("failed to handle Yookassa webhook: %w", err)
	}
	return nil
}
