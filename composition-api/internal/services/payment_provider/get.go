package payment_provider

import (
	"context"

	domain "composition-api/internal/domain/billing"
)

func (s *service) ListPaymentProviders(ctx context.Context) ([]domain.PaymentProvider, error) {
	return s.adapters.Billing.ListPaymentProviders(ctx)
}
