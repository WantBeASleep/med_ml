package yookassa_webhook

import (
	"context"

	domain "composition-api/internal/domain/billing"
)

func (s *service) HandleYookassaWebhook(ctx context.Context, req domain.YookassaWebhookRequest) error {
	return s.adapters.Billing.HandleYookassaWebhook(ctx, req)
}
