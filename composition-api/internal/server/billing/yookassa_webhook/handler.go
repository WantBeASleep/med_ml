package yookassa_webhook

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type YookassaWebhookHandler interface {
	YookassaWebhooksPost(ctx context.Context, req *api.YookassaWebhookRequest) (api.YookassaWebhooksPostRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) YookassaWebhookHandler {
	return &handler{
		services: services,
	}
}
