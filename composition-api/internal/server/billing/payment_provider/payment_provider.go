package payment_provider

import (
	"context"

	api "composition-api/internal/generated/http/api"

	"github.com/AlekSi/pointer"
)

func (h *handler) PaymentProvidersGet(ctx context.Context) (api.PaymentProvidersGetRes, error) {
	providers, err := h.services.PaymentProviderService.ListPaymentProviders(ctx)
	if err != nil {
		return nil, err
	}

	var apiProviders []api.PaymentProvider
	for _, provider := range providers {
		apiProviders = append(apiProviders, api.PaymentProvider{
			ID:       api.NewOptUUID(provider.ID),
			Name:     api.NewOptString(provider.Name),
			IsActive: api.NewOptBool(provider.IsActive),
		})
	}

	return pointer.To(api.PaymentProvidersGetOKApplicationJSON(apiProviders)), nil
}
