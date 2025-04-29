package billing

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	domain "composition-api/internal/domain/billing"

	"github.com/google/uuid"
)

func (a *adapter) ListPaymentProviders(ctx context.Context) ([]domain.PaymentProvider, error) {
	res, err := a.client.ListPaymentProviders(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to list payment providers: %w", err)
	}

	var providers []domain.PaymentProvider
	for _, provider := range res.PaymentProviders {
		providers = append(providers, domain.PaymentProvider{
			ID:       uuid.MustParse(provider.Id),
			Name:     provider.Name,
			IsActive: provider.IsActive,
		})
	}
	return providers, nil
}
