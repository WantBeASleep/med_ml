package payment_provider

import (
	"context"

	"composition-api/internal/adapters"
	domain "composition-api/internal/domain/billing"
)

type Service interface {
	ListPaymentProviders(ctx context.Context) ([]domain.PaymentProvider, error)
}

type service struct {
	adapters *adapters.Adapters
}

func New(adapters *adapters.Adapters) Service {
	return &service{
		adapters: adapters,
	}
}
