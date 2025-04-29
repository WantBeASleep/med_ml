package yookassa_webhook

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	domain "composition-api/internal/domain/billing"
	"composition-api/internal/generated/http/api"

	"github.com/AlekSi/pointer"
)

func mapYookassaStatusToPaymentStatus(yookassaStatus string) (domain.PaymentStatus, error) {
	switch yookassaStatus {
	case "pending":
		return domain.PayPending, nil
	case "waiting_for_capture":
		return domain.PayWaitingForCapture, nil
	case "succeeded":
		return domain.PayCompleted, nil
	case "canceled":
		return domain.PayCancelled, nil
	default:
		return "", fmt.Errorf("unknown Yookassa status: %s", yookassaStatus)
	}
}

func (h *handler) YookassaWebhooksPost(ctx context.Context, req *api.YookassaWebhookRequest) (api.YookassaWebhooksPostRes, error) {
	event, _ := req.Event.Get()
	paymentID := req.Object.Value.ID.Value
	yookassaStatus := req.Object.Value.Status.Value
	status, err := mapYookassaStatusToPaymentStatus(yookassaStatus)
	if err != nil {
		return nil, err
	}
	slog.Info(string(status))

	receivedAt := time.Now()

	extraInfoBytes, err := req.MarshalJSON()
	if err != nil {
		return nil, err
	}
	extraInfoJSON := string(extraInfoBytes)

	extraInfo := map[string]string{
		"json": extraInfoJSON,
	}

	webhookRequest := domain.YookassaWebhookRequest{
		Event:        event,
		PaymentID:    paymentID.String(),
		Status:       string(status),
		ProviderName: "Yookassa",
		ReceivedAt:   receivedAt,
		ExtraInfo:    extraInfo,
		Reason:       "",
	}

	err = h.services.YookassaWebhookService.HandleYookassaWebhook(ctx, webhookRequest)
	if err != nil {
		return nil, err
	}

	return pointer.To(api.YookassaWebhooksPostOK{}), nil
}
