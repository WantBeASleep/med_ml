package payment_provider

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type PaymentProviderHandler interface {
	PaymentProvidersGet(ctx context.Context) (api.PaymentProvidersGetRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) PaymentProviderHandler {
	return &handler{
		services: services,
	}
}
