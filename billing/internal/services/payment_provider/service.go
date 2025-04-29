package payment_provider

import (
	"context"
	"fmt"

	"billing/internal/domain"
	"billing/internal/repository"
)

type Service interface {
	ListPaymentProviders(ctx context.Context) ([]domain.PaymentProvider, error)
}

type service struct {
	dao repository.DAO
}

func New(dao repository.DAO) Service {
	return &service{
		dao: dao,
	}
}

func (s *service) ListPaymentProviders(ctx context.Context) ([]domain.PaymentProvider, error) {
	paymentProvidersDB, err := s.dao.NewPaymentProviderQuery(ctx).ListPaymentProviders()
	if err != nil {
		return nil, fmt.Errorf("list payment providers: %w", err)
	}

	var paymentProviders []domain.PaymentProvider
	for _, p := range paymentProvidersDB {
		paymentProviders = append(paymentProviders, p.ToDomain())
	}

	return paymentProviders, nil
}
